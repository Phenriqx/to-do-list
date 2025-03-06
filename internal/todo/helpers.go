package todo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
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