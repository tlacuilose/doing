package store

import (
	"testing"

	"github.com/tlacuilose/doing/task"
)

const testFile string = "./test_store.txt"

// Test saving a task in a file store.
// Is ran first to create the file if it doesnt exist.
func TestSaveInStore(t *testing.T) {
	store := New(testFile)

	tasksToSave := []task.Task{task.New("New Task")}
	err := store.SaveTasks(tasksToSave)
	if err != nil {
		t.Fatal(err)
	}
}

// Test reading a task in a file store.
func TestReadFromStore(t *testing.T) {
	store := New(testFile)

	_, err := store.ReadTasks()
	if err != nil {
		t.Fatal(err)
	}
}

// Test savind  a task in a file store and.
// Test reading a task in a file store.
func TestCorrectSaveAndReadStore(t *testing.T) {
	store := New(testFile)

	tasksToSave := []task.Task{task.New("New Task"), task.New("Second Task")}
	err := store.SaveTasks(tasksToSave)
	if err != nil {
		t.Fatal(err)
	}

	savedTasks, err := store.ReadTasks()
	if err != nil {
		t.Fatal(err)
	}

	if (savedTasks)[0].Description != tasksToSave[0].Description {
		t.Fatalf(`Couldn't save task "%s" in store.`, tasksToSave[0].Description)
	}
}
