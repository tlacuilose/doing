// Package stores tasks in a text file.
// First define the Store with New(filename)
package store

import (
	"bufio"
	"os"

	"github.com/tlacuilose/doing/task"
)

// Store is passed as the place to save tasks.
type Store struct {
	fileName string
}

// New creates a store to pass on other functions in this package.
// Is defined by a file name for the text file where tasks are stored.
func New(fileName string) Store {
	return Store{fileName}
}

// ReadTasks reads the tasks onto a Task slice from a Store.
func (s Store) ReadTasks() ([]task.Task, error) {
	tasks := make([]task.Task, 0)

	file, err := os.Open(s.fileName)
	if err != nil {
		return tasks, err
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		line := sc.Text()
		task, err := task.Decode(line)
		if err != nil {
			continue
		}

		tasks = append(tasks, *task)

	}

	return tasks, nil
}

// SaveTasks saves the tasks from a Task slice onto a Store.
func (s Store) SaveTasks(tasks []task.Task) error {
	file, err := os.Create(s.fileName)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		encoded, err := task.Encode(&t)
		if err != nil {
			continue
		}
		_, err = file.WriteString(encoded)
		if err != nil {
			continue
		}
	}
	return nil
}
