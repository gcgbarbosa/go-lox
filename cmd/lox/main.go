package main

import (
	"fmt"
	"os"

	"gcgbarbosa.com/lox/pkg/lox"
)

func main() {
	// Check if a parameter was provided
	status := lox.Lox(os.Args)

	fmt.Print("Exiting with status:", status)
}
