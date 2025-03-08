package todo

import (
	"errors"
	"fmt"
	"io"

	// "io"
	"log"
	"os"

	// "bufio"
	"encoding/json"
)

func LoadTodos(filename string) ([]TodoItem, error) {
	var todos []TodoItem

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// if the file does not exist, create an empty slice
			return []TodoItem{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []TodoItem{}, nil
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		if todo.ID > lastID {
			lastID = todo.ID
		}
	}
	// fmt.Println("Loading todos from file...")
	return todos, nil
}

func SaveTodo(filename string, todos []TodoItem) error {
	jsonData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	fmt.Println("Saving todos to file...")
	return os.WriteFile(filename, jsonData, 0644)
}

func CreateTodo(title string, description string) TodoItem {
	lastID++
	return TodoItem{
		ID:          lastID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
}

func GetTodos(filename string) ([]TodoItem, error) {
	var todos []TodoItem
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return todos, json.Unmarshal(fileData, &todos)
}

func GetTodoByID(todoMap map[uint]TodoItem, id uint) (TodoItem, error) {
	// fmt.Println(todoMap)
	todo, exists := todoMap[id]
	if !exists {
		return TodoItem{}, errors.New("Todo not found")
	}
	return todo, nil
}

func UpdateTodo(filename string, id uint, title, description string, completed bool) (TodoItem, error) {
	var todos []TodoItem
	var UpdatedTodo TodoItem
	jsonFile, err := os.Open(filename)
	if err != nil {
		return UpdatedTodo, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return UpdatedTodo, err
	}
	if err := json.Unmarshal(byteValue, &todos); err != nil {
		return UpdatedTodo, err
	}

	found := false
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = title
			todos[i].Description = description
			todos[i].Completed = completed
			UpdatedTodo = todos[i]
			found = true
			break
		}
	}
	if !found {
		return UpdatedTodo, errors.New("Todo not found")
	}

	// Save the updated todos to the file
	jsonFile, err = os.Create(filename)
	if err != nil {
		return UpdatedTodo, err
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(todos)
	if err != nil {
		return UpdatedTodo, err
	}	
	return UpdatedTodo, nil
}