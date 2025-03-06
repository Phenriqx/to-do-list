package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/phenriqx/todo-list-golang/internal/todo"
)

func main() {
	commands := []string {
		"create",
		"get",
        "update",
        "delete",
	}

	fmt.Println("Welcome to the TODO App!")
	fmt.Println("Type one of the following commands:")

	for index, cmd := range commands {
        fmt.Printf("%d - %s\n", index + 1, cmd)
    }

	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading command:", err)
		os.Exit(1)
	} else if command == "" || len(command) <= 1 {
		fmt.Println("Invalid command. Please try again.")
		os.Exit(1)
	}

	switch {
		case command == "create\n":
            todo := todo.CreateTodo()
            fmt.Printf("Created todo with ID: %d\n", todo.ID)
        case command == "get\n":
            fmt.Println("Get command not implemented yet.")
        case command == "update\n":
            fmt.Println("Update command not implemented yet.")
        case command == "delete\n":
            fmt.Println("Delete command not implemented yet.")
        default:
            fmt.Printf("Invalid command: %s. Please try again.\n", command)
            os.Exit(1)
	}
}
