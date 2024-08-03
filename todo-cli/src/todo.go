package main

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

type Todo struct {
	ID   int
	Name string
	Done bool
}

func addNewTodo(todos []Todo) []Todo {
	name := readInput("Enter what you want to do:")

	newTodo := Todo{
		ID:   rand.IntN(100000),
		Name: name,
		Done: false,
	}

	todos = append(todos, newTodo)

	return todos
}

func markAsDone(todos []Todo) []Todo {
	id, err := chooseTodo(todos)

	if err != nil {
		return todos
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			break
		}
	}

	return todos
}

func removeTodo(todos []Todo) []Todo {
	id, err := chooseTodo(todos)

	if err != nil {
		return todos
	}

	newTodos := []Todo{}

	for _, todo := range todos {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		}
	}

	return newTodos
}

func chooseTodo(todos []Todo) (int, error) {
	var selectedTodo string

	options := make([]string, 0, len(todos))

	for _, todo := range todos {
		options = append(options, fmt.Sprintf("#%v %v", todo.ID, todo.Name))
	}

	prompt := &survey.Select{
		Message: "Which one?",
		Options: options,
	}

	survey.AskOne(prompt, &selectedTodo)

	re := regexp.MustCompile(`#(\d+)`)
	matches := re.FindStringSubmatch(selectedTodo)

	return strconv.Atoi(matches[1])
}

func printTodos(todos []Todo) {
	fmt.Println("TODO List:")
	for _, todo := range todos {
		status := "Pending"

		if todo.Done {
			status = "Done"
		}

		fmt.Printf("ID: %d, Title: %s, Status: %s\n", todo.ID, todo.Name, status)
	}
}
