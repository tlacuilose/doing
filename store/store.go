// Package stores tasks in a text file.
// First define the Store with New(filename)
package store

import "github.com/tlacuilose/doing/task"

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
func (s Store) ReadTasks(tasks *[]task.Task) error {
	return nil
}

// SaveTasks saves the tasks from a Task slice onto a Store.
func (s Store) SaveTasks(tasks *[]task.Task) error {
	return nil
}
