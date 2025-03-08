package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var todos []TodoItem
	if err := json.Unmarshal(byteValue, &todos); err != nil {
		return nil, err
	}

	// Convert slice to a map for O(1) lookups
	todoMap := make(map[uint]TodoItem)
	for _, todo := range todos {
		todoMap[todo.ID] = todo
	}
	return todoMap, nil
}

func ToInt (s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ToBool (s string) (bool, error) {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return i, nil
}