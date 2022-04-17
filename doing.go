package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/tlacuilose/doing/command"
	"github.com/tlacuilose/doing/store"
	"github.com/tlacuilose/doing/task"
)

// Default store location.
var defaultFileName string = "./store/store.txt"
var s store.Store = store.New(defaultFileName)

func listAllTasks() error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return errors.New("No saved tasks to read.")
	}

	for i, t := range tasks {
		fmt.Printf("%d ", i)
		fmt.Println(t.Print())
	}

	return nil
}

func listDoingTasks() error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return errors.New("No saved tasks to read.")
	}

	for _, t := range tasks {
		if t.Status == task.Doing {
			fmt.Println(t.Print())
		}
	}

	return nil
}

func listCompletedTasks() error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return errors.New("No saved tasks to read.")
	}

	for _, t := range tasks {
		if t.Status == task.Completed {
			fmt.Println(t.Print())
		}
	}

	return nil
}

func addTask(description string) error {
	tasks, _ := s.ReadTasks() // Failing to read returns an empty task list.

	newTask := task.New(description)
	tasks = append([]task.Task{newTask}, tasks...)

	fmt.Printf("Added new task!\n%s\n", newTask.Print())

	return s.SaveTasks(tasks)
}

func completeTask(reverseIndex int) error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return err
	}

	if reverseIndex < 0 || reverseIndex >= len(tasks) {
		return errors.New("Please provide a valid index for a task.")
	}

	tasks[reverseIndex].Complete()

	fmt.Printf("Completed task!\n%s\n", tasks[reverseIndex].Print())
	return s.SaveTasks(tasks)
}

func deleteTask(reverseIndex int) error {
	tasks, err := s.ReadTasks()
	if err != nil {
		return err
	}

	if reverseIndex < 0 || reverseIndex >= len(tasks) {
		return errors.New("Please provide a valid index for a task.")
	}

	tasks = append(tasks[:reverseIndex], tasks[reverseIndex+1:]...)
	err = s.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Deleted task. All tasks:")
	return listAllTasks()
}

func resetTasks() error {
	tasks := make([]task.Task, 0)
	err := s.SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Reset store, deleted all saved tasks.")
	return nil
}

func main() {
	cmd, err := command.ParseFromShell()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch cmd.Type {
	case command.ListAll:
		err = listAllTasks()
	case command.ListDoing:
		err = listDoingTasks()
	case command.ListCompleted:
		err = listCompletedTasks()
	case command.AddTask:
		err = addTask(cmd.Body)
	case command.CompleteTask:
		reverseIndex, _ := strconv.Atoi(cmd.Body)
		err = completeTask(reverseIndex)
	case command.DeleteTask:
		reverseIndex, _ := strconv.Atoi(cmd.Body)
		err = deleteTask(reverseIndex)
	case command.ResetTasks:
		err = resetTasks()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
