package task

import (
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
