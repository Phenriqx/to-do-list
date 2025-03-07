package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"io"
)

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatal("Invalid input. Please try again.")
	}
	return input
}

// GetTodoMap loads the todos from a JSON file and returns them as a map for faster lookups.
func GetTodoMap(filename string) (map[uint]TodoItem, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer jsonFile.Close()
	byteValye, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var todos []TodoItem
	if err := json.Unmarshal(byteValye, &todos); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Convert slice to a map for O(1) lookups
	todoMap := make(map[uint]TodoItem)
	for _, todo := range todos {
		todoMap[todo.ID] = todo
	}
	return todoMap, nil
}
