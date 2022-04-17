package task

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

// Test the creation of a new task with the current time and a status of doing.
func TestNewTask(t *testing.T) {
	now := time.Now().Format(time.RFC850)
	expectedTask := "new task"
	initialStatus := Doing
	task := New(expectedTask)

	if task.CreationTime != now {
		t.Fatal("Task time is incorrect")
	}

	if task.Description != expectedTask {
		t.Fatal("Task task is incorrect")
	}

	if task.Status != initialStatus {
		t.Fatal("Task status is incorrect")
	}
}

// Test that a new task can be marked as completed.
func TestMarkCompletedTask(t *testing.T) {
	now := time.Now().Format(time.RFC850)
	expectedTask := "A new task"
	task := New(expectedTask)
	task.Complete()

	if task.CreationTime != now {
		t.Fatal("Task time was changed when marked as completed.")
	}

	if task.Description != expectedTask {
		t.Fatal("Task description was changed when marked as completed.")
	}

	if task.Status != Completed {
		t.Fatal("Task could not be marked as completed.")
	}
}

// Test printing a task with the current time and doing status and then completed.
func TestPrintDoingTask(t *testing.T) {
	now := time.Now().Format(time.RFC850)
	expectedTask := "A new task"
	task := New(expectedTask)
	expectedDoingPrint := fmt.Sprintf("[Doing] %s created: %s\n", expectedTask, now)

	if task.Print() != expectedDoingPrint {
		t.Fatal("Doing task was incorrectly printed.")
	}

	task.Complete()
	expectedCompletedPrint := fmt.Sprintf("[Completed] %s created: %s\n", expectedTask, now)

	if task.Print() != expectedCompletedPrint {
		t.Fatal("Completed task was incorrectly printed.")
	}
}

// Test encoding a task into a b64,b64,int.
func TestEncodeTask(t *testing.T) {
	timeString := "Friday, 15-Apr-22 19:25:14 CDT"
	btime := base64.StdEncoding.EncodeToString([]byte(timeString))

	descriptionString := "A new task"
	bdescription := base64.StdEncoding.EncodeToString([]byte(descriptionString))

	var status int64 = 0

	expectedEncodedTask := fmt.Sprintf("%s,%s,%d\n", btime, bdescription, status)

	task := Task{timeString, descriptionString, Status(status)}

	encodedTask, err := Encode(&task)
	if err != nil {
		t.Fatal(err)
	}

	if encodedTask != expectedEncodedTask {
		t.Fatal("Failed to encode task.")
	}

}

// Test decoding a task from a b64,b64,int to a task.
func TestDecodeTask(t *testing.T) {
	timeString := "Friday, 15-Apr-22 19:25:14 CDT"
	btime := base64.StdEncoding.EncodeToString([]byte(timeString))

	descriptionString := "A new task"
	bdescription := base64.StdEncoding.EncodeToString([]byte(descriptionString))

	status := Status(0)

	// No \n because scan() should read a line.
	encodedTask := fmt.Sprintf("%s,%s,%d", btime, bdescription, status)

	task, err := Decode(encodedTask)
	if err != nil {
		t.Fatal(err)
	}

	if task.CreationTime != timeString {
		t.Fatal("Couldnt get creation time of task.")
	}

	if task.Description != descriptionString {
		t.Fatal("Couldnt get description of task.")
	}

	if task.Status != status {
		t.Fatal("Couldnt get status of task.")
	}
}
