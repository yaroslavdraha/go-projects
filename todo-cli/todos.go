package main

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strconv"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

type Todo struct {
	ID   int
	Name string
	Done bool
}

func (todo Todo) printTodo() {
	fmt.Printf("qwe %v", todo.Name)
}

func addNewTodo(todos []Todo) []Todo {
	name := readInput("Enter what you want to do:")

	newTodo := Todo{
		ID:   rand.IntN(100000),
		Name: name,
		Done: false,
	}

	todos = append(todos, newTodo)

	wg.Add(1)
	go saveTodoInDB(newTodo)

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

func saveTodoInDB(todo Todo) {
	time.Sleep(10 * time.Second)

	fmt.Println("Todo was saved to DB")

	wg.Done()
}
