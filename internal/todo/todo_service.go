package todo

import (
	"fmt"
	"os"

	//"encoding/json"
	"bufio"
)

func CreateTodo() (TodoItem) {
	lastID++

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your todo title:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading title:", err)
        os.Exit(1)
    } else if title == "" || len(title) <= 1 {
		fmt.Println("Invalid title. Please try again.")
	    os.Exit(1)
	}

	fmt.Println("Enter your todo description:")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading description:", err)
		os.Exit(1)
	} else if description == "" || len(description) <= 1 {
		fmt.Println("Invalid description. Please try again.")
        os.Exit(1)
	}

	todo := TodoItem {
		ID: lastID,
        Title: title,
        Description: description,
        Completed: false,
	}
	return todo
}
