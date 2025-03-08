package main

import (
	"bufio"
	"bytes"
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
	fmt.Println()

	for index, cmd := range commands {
		fmt.Printf("%d - %s\n", index+1, cmd)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	command := todo.GetInput("Enter a command: ")

	switch {
	case command == "create" || command == "1":

		title := todo.GetInput("Enter the title of the todo:")
		description := todo.GetInput("Enter the description of the todo:")

		newTodo := todo.CreateTodo(title, description)
		todos = append(todos, newTodo)
		err := todo.SaveTodo(filename, todos)
		if err != nil {
			log.Fatal("Error saving todos:", err)
		}

		fmt.Println("Created todo with ID:", newTodo.ID)

	case command == "get" || command == "2":

		todo_items, err := todo.GetTodos(filename)
		if err != nil {
			log.Fatal("Error getting todos:", err)
		}
		// Formmat the output for better readability
		var buffer bytes.Buffer
		for _, todo := range todo_items {
			buffer.WriteString(fmt.Sprintf("ID: %d,\n Title: %s,\n Description: %s,\n Completed: %t\n", 
			todo.ID, todo.Title, todo.Description, todo.Completed))
		}
		fmt.Println(buffer.String())

	case command == "get by ID" || command == "3":

		todoMap, err := todo.GetTodoMap(filename)
		if err != nil {
			log.Fatal("Error getting todo map:", err)
		}

		idStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading ID:", err)
		}
		idStr = idStr[:len(idStr)-1] // Remove the newline character
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("Error parsing ID:", err)
		}
		todo, err := todo.GetTodoByID(todoMap, uint(id))
		if err != nil {
			log.Fatal("Error getting todo by ID:", err)
		}

		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("ID: %d,\n Title: %s,\n Description: %s,\n Completed: %t\n",
				todo.ID, todo.Title, todo.Description, todo.Completed))

		fmt.Println(buffer.String())


	case command == "update" || command == "4":

		idStr := todo.GetInput("Enter the ID of the todo to update: ")
		id, err := todo.ToInt(idStr)
		if err != nil {
			log.Fatal("Error parsing ID:", err)
		}
		lastID := todo.GetLastID(todos)
		if err := todo.CheckID(uint(id), lastID); err != nil {
			log.Fatal("Index out of range!")
		}

		title := todo.GetInput("Enter the new title (leave blank to keep the same): ")
		description := todo.GetInput("Enter the new description (leave blank to keep the same): ")
		completedStr := todo.GetInput("Enter the new completed status (true/false) (leave blank to keep the same): ")
		completedBool, err := todo.ToBool(completedStr)
		if err != nil {
			log.Fatal("Error parsing completed status:", err)
		}

		todo, err := todo.UpdateTodo(filename, uint(id),
			title, description, completedBool)
		if err != nil {
			log.Fatal("Error updating todo:", err)
		}
		
		fmt.Println("Updated todo with ID:", todo.ID)

	case command == "delete" || command == "5":
		
		idStr := todo.GetInput("Enter the ID of the todo to delete: ")
		id, err := todo.ToInt(idStr)
		if err != nil {
			log.Fatal("Error parsing ID:", err)
		}

		err = todo.DeleteTodo(filename, uint(id))
		if err != nil {
            log.Fatal("Error deleting todo:", err)
        }

	default:
		fmt.Printf("Invalid command: %s. Please try again.\n", command)
		os.Exit(1)
	}
}