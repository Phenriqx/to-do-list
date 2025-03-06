package main

import (
	"bufio"
	"log"
	// "errors"
	"fmt"
	"os"
	"strconv"

	"github.com/phenriqx/todo-list-golang/internal/todo"
)

func main() {
	filename := "data/todos.json"
	todos, err := todo.LoadTodos(filename)
	if err != nil {
        log.Fatal("Error loading todos:", err)
    }

	commands := []string{
		"create",
		"get",
		"get by ID",
		"update",
		"delete",
	}

	fmt.Println("Welcome to the TODO App!")
	fmt.Println("Type one of the following commands:")

	for index, cmd := range commands {
		fmt.Printf("%d - %s\n", index+1, cmd)
	}

	reader := bufio.NewReader(os.Stdin)
	command := todo.GetInput("Enter a command:")

	switch {
	case command == "create":
		title := todo.GetInput("Enter the title of the todo:")
		description := todo.GetInput("Enter the description of the todo:")

		todo_item := todo.CreateTodo(title, description)
		todos = append(todos, todo_item)
		err := todo.SaveTodo(filename, todos)
		if err != nil {
			log.Fatal("Error saving todos:", err)
		}

		fmt.Println("Created todo with ID:", todo_item.ID)

	case command == "get\n":
		todo, err := todo.GetTodos()
		if err != nil {
            log.Fatal("Error getting todos:", err)
        }
		fmt.Println(todo)

	case command == "get by ID\n":
		idStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading ID:", err)
		}
		idStr = idStr[:len(idStr)-1] // Remove the newline character
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("Error parsing ID:", err)
		}
		todo, err := todo.GetTodoByID(uint(id))
		if err != nil {
			log.Fatal("Error getting todo by ID:", err)
		}
		fmt.Println(todo)

	case command == "update\n":
		fmt.Println("Update command not implemented yet.")
	case command == "delete\n":
		fmt.Println("Delete command not implemented yet.")
	default:
		fmt.Printf("Invalid command: %s. Please try again.\n", command)
		os.Exit(1)
	}
}

// func toInt(s string) (uint, error) {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, errors.New("Invalid input. Please enter a valid number.")
// 	}
// 	return uint(i), nil
// }
