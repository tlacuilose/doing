// Package defines the task model.
package task

import "time"

// Status is the status of the task.
type Status int64

// All available status for a task.
const (
	Doing Status = iota
	Completed
)

// Task has time when it was created, a description and a status (doing or completed)
type Task struct {
	CreationTime string
	Description  string
	Status       Status
}

// New creates a new task with the current time and a status of doing.
func New(description string) Task {
	return Task{
		CreationTime: time.Now().Format(time.RFC850),
		Description:  description,
		Status:       Doing,
	}
}
