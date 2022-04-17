// Package defines the task model.
package task

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Status is the status of the task.
type Status int64

// All available status for a task.
const (
	Doing Status = iota
	Completed
)

func getStatus(s Status) string {
	switch s {
	case Doing:
		return "[Doing]\t"
	case Completed:
		return "[Completed]"
	}
	return "404"
}

// Group all status for reference of length and members.
var allStatus = []Status{Doing, Completed}

// Task has time when it was created, a description and a status (doing or completed)
type Task struct {
	CreationTime string
	Description  string
	Status       Status
}

// Validates that the task has full information
func (t *Task) validate() error {
	if t.CreationTime == "" || t.Description == "" || t.Status < 0 {
		return errors.New("Malformed task.")
	}
	return nil
}

// Print a task in format [Status] Description created: CreationTime
func (t *Task) Print() string {
	return fmt.Sprintf("(%s) %s\t %s", t.CreationTime, getStatus(t.Status), t.Description)
}

// Complete, change task status to completed.
func (t *Task) Complete() {
	t.Status = Completed
}

// New creates a new task with the current time and a status of doing.
func New(description string) Task {
	return Task{
		CreationTime: time.Now().Format(time.RFC822),
		Description:  description,
		Status:       Doing,
	}
}

// Encodes a string to bs64 string.
func b64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// Decodes a string to bs64 string.
func b64Decode(encodedText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodedText)
	return string(data), err
}

// Encodes task into base64 single line.
// An encoded task is of string function "base64(CreationTime),base64(Description),Status
func Encode(task *Task) (string, error) {
	if err := task.validate(); err != nil {
		return "", err
	}

	bTime := b64Encode(task.CreationTime)
	bDescription := b64Encode(task.Description)

	encodedTask := fmt.Sprintf("%s,%s,%d\n", bTime, bDescription, task.Status)

	return encodedTask, nil
}

// Decode task from a string line.
// An encoded task is of string function "base64(CreationTime),base64(Description),Status
func Decode(encodedText string) (*Task, error) {
	fields := strings.Split(encodedText, ",")
	if len(fields) < 3 {
		return &Task{}, errors.New("Could not get all field from encoded task.")
	}

	time, err := b64Decode(fields[0])
	if err != nil {
		return &Task{}, err
	}

	description, err := b64Decode(fields[1])
	if err != nil {
		return &Task{}, err
	}

	statusCode, err := strconv.Atoi(fields[2])
	if err != nil {
		return &Task{}, err
	}

	if statusCode > int(allStatus[len(allStatus)-1]) {
		return &Task{}, errors.New("Could not get a valid status from encoded task.")
	}

	status := Status(statusCode)

	return &Task{time, description, status}, nil
}
