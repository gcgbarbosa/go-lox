package lox

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hadError bool = false

func report(line int, where string, message string) {
	fmt.Println("[line " + string(line) + "] Error" + where + ": " + message)
	hadError = true
}

func error(line int, message string) {
	report(line, "", message)
}

func run(line string) {
	scanner := bufio.NewScanner(strings.NewReader(line))

	// Define the scanner to split input by spaces (or other delimiter)
	scanner.Split(bufio.ScanWords)

	// Scan through each token (word in this case)
	for scanner.Scan() {
		token := scanner.Text() // Get the current token
		fmt.Println("Scanned Token:", token)
	}
}

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

		fmt.Println("Running: <" + input + ">")
		run(input)
	}
}

func Lox(args []string) int {
	// don't run code with errors
	if hadError {
		return 65
	}

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
