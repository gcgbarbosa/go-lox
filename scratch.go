package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runFile(path string) {
	fmt.Println("Running file <" + path + ">")
}

func runPrompt() {
	fmt.Println("Running prompt")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		fmt.Println("---")
		fmt.Println("Running: <" + input + ">")
	}
}

func lox(args []string) int {
	if len(args) > 2 {
		fmt.Println("Usage: jlox [script]")
		return 64
	} else if len(args) == 2 {
		param := args[1]
		runFile(param)
	} else {
		runPrompt()
	}

	fmt.Println(args)
	return 0
}

func main() {
	// Check if a parameter was provided
	status := lox(os.Args)

	fmt.Print("Exiting with status:", status)
}
