package main

import "os"

type Todo struct {
	Time   string
	Task   string
	Status string // TODO: enumeration:?
}

func parseTodoFromString(line string) (Todo, error) {
	return Todo{}, nil
}

func getTodosFromFile(file os.File) ([]Todo, error) {
	return make([]Todo, 0), nil
}

func addTodoToFile(file os.File) error {
	return nil
}

func printTodos(todos []Todo) error {
	return nil
}

func printCompletedTodos(todos []Todo) error {
	return nil
}

func printUncompletedTodos(todos []Todo) error {

}

func main() {

}
