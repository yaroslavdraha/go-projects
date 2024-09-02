package main

import (
	"fmt"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

var wg = sync.WaitGroup{}

var actionOptions = []string{SHOW_ALL, ADD_NEW_TODO, MARK_AS_DONE, REMOVE_TODO, EXIT}

func main() {
	fmt.Println("Welcome to CLI TODOs")

	todos := []Todo{}

	var test = make(map[string]string)
	test["name"] = "Clean room"
	test["isDone"] = "done"

	// loop:
	// 	for {
	isEmpty := !toBool(len(todos))

	switch action := askAction(isEmpty); action {
	case SHOW_ALL:
		printTodos(todos)
	case ADD_NEW_TODO:
		todos = addNewTodo(todos)
		fmt.Printf("You've added todo with ID #%v\n", todos[len(todos)-1].ID)
	case MARK_AS_DONE:
		todos = markAsDone(todos)
	case REMOVE_TODO:
		todos = removeTodo(todos)
	case EXIT:
		// break loop
	}

	wg.Wait()

	// }
}

func askAction(isEmpty bool) string {
	var action string

	actions := make([]string, len(actionOptions))
	copy(actions, actionOptions)

	if isEmpty {
		actions = actions[1:]
	}

	prompt := &survey.Select{
		Message: "What would you like to do?",
		Options: actions,
	}

	survey.AskOne(prompt, &action)

	return action
}
