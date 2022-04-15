// Package finds the expected command to be run.
package command

import (
	"errors"
	"flag"
	"strconv"
)

// CommandType describes all the available commands for this tool.
type CommandType int64

// All available commands.
const (
	ListAll CommandType = iota
	ListDoing
	ListCompleted
	AddTask
	CompleteTask
	DeleteTask
	ResetTasks
)

// Command is the returned struct with the command type.
// The body holds any text after the command flag definition.
type Command struct {
	Type CommandType
	Body string
}

// ParseFlagsAndArgs parses the flags from running the program.
// Returns a Command with the command type or an error.
func parseFlagsAndArgs() (Command, error) {

	listAllFlag := flag.Bool("la", false, "List all tasks: doing and completed.")
	listDoingFlag := flag.Bool("ld", false, "List all tasks doing")
	listCompletedFlag := flag.Bool("lc", false, "List all tasks completed")
	addTaskFlag := flag.Bool("a", false, "Add new task")
	completeTaskFlag := flag.Bool("c", false, "Mark task as completed")
	deleteTaskFlag := flag.Bool("d", false, "Delete task")
	resetTasksFlag := flag.Bool("r", false, "Reset all tasks, permanently deletes all tasks.")
	flag.Parse()

	// First flag recognized in the above order is returned.

	isList := *listAllFlag || *listDoingFlag || *listCompletedFlag

	if isList {
		switch {
		case *listAllFlag:
			return Command{ListAll, ""}, nil
		case *listDoingFlag:
			return Command{ListDoing, ""}, nil
		case *listCompletedFlag:
			return Command{ListCompleted, ""}, nil
		}
	}

	isAction := *addTaskFlag || *completeTaskFlag || *deleteTaskFlag

	// Evaluate reset tasks flags last
	if !isAction {
		if *resetTasksFlag {
			return Command{ResetTasks, ""}, nil
		} else {
			return Command{}, errors.New("Please provide a valid command.")
		}
	}

	// Actions require a command body.

	args := flag.Args()
	if len(args) != 1 {
		return Command{}, errors.New("Please provide a valid command.")
	}

	commandBody := args[0]

	if isAction {
		switch {
		case *addTaskFlag:
			return Command{AddTask, commandBody}, nil
		case *completeTaskFlag:
			return Command{CompleteTask, commandBody}, nil
		case *deleteTaskFlag:
			return Command{DeleteTask, commandBody}, nil
		}
	}

	return Command{}, errors.New("Please provide a valid command.")
}

func validateCommandBody(command Command) error {
	if command.Type == CompleteTask || command.Type == DeleteTask {
		_, err := strconv.Atoi(command.Body)
		if err != nil {
			return errors.New("An integer id is needed for marking task as completed or deleted.")
		}
	}
	return nil
}

// ParseFromShell parses the desired command from the shell input.
// Returns a Command with type and body or an error.
func ParseFromShell() (Command, error) {
	cmd, err := parseFlagsAndArgs()
	if err != nil {
		return cmd, err
	}
	err = validateCommandBody(cmd)
	return cmd, err
}
