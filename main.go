package main

import (
	"fmt"

	"gcgbarbosa.com/lox/lox"
)

func main() {
	token := lox.Token{
		Literal: "test",         // No literal value
		Lexeme:  "(",            // Lexeme representing the token
		Line:    1,              // Line where the token is found
		Type:    lox.LESS_EQUAL, // Type of token
	}

	fmt.Println(lox.ToString(token))
	fmt.Println(token)

	k := lox.ScanTokens("()")

	fmt.Println(*k)

	// Check if a parameter was provided
	// status := lox.Lox(os.Args)
	// fmt.Print("Exiting with status:", status)
}
