package store

import (
	"testing"

	"github.com/tlacuilose/doing/task"
)

func TestReadFromStore(t *testing.T) {
	store := New("any_file")

	savedTasks := make([]task.Task, 0)
	err := store.ReadTasks(&savedTasks)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveInStore(t *testing.T) {
	store := New("any_file")

	tasksToSave := []task.Task{task.New("New Task")}
	err := store.SaveTasks(&tasksToSave)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCorrectSaveAndReadStore(t *testing.T) {
	store := New("any_file")

	tasksToSave := []task.Task{task.New("New Task")}
	err := store.SaveTasks(&tasksToSave)
	if err != nil {
		t.Fatal(err)
	}

	savedTasks := make([]task.Task, 0)
	err = store.ReadTasks(&savedTasks)
	if err != nil {
		t.Fatal(err)
	}

	if savedTasks[0].Description != tasksToSave[0].Description {
		t.Fatalf(`Couldn't save task "%s" in store.`, tasksToSave[0].Description)
	}
}
