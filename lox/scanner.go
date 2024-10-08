package lox

import "fmt"

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string, tokens []Token) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  tokens,
		start:   0,
		current: 0,
		line:    1,
	}
}

func ScanTokens(source string) *Scanner {
	// scanner := NewScanner(source)
	tokens := make([]Token, 0)

	start := 0
	current := 0
	line := 1

	isAtEnd := func() bool {
		return current >= len(source)
	}

	advance := func() rune {
		current += 1

		runes := []rune(source)

		fmt.Println("Current:", current)

		return runes[current-1]
	}

	addTokenAndLiteral := func(type_ TokenType, literal interface{}) {
		text := string(source[start:current])
		tokens = append(tokens, Token{Type: type_, Lexeme: text, Literal: literal, Line: line})
	}

	// Helper function that always passes nil for the literal
	addToken := func(type_ TokenType) {
		addTokenAndLiteral(type_, nil)
	}

	// TODO: add missing tokens
	scanToken := func() {
		c := advance()

		switch c {
		case '(':
			addToken(LEFT_PAREN)
		case ')':
			addToken(RIGHT_PAREN)
		// Add more cases for other tokens like operators, literals, etc.
		default:
			// Handle unknown characters or skip whitespace here
			fmt.Println("Unknown character:", c)
		}
	}

	for !isAtEnd() { // while !isAtEnd()
		// if isAtEnd() {
		// 	break
		// }

		scanToken()

		start = current
	}

	fmt.Println(tokens)

	// add EOF to end read
	tokens = append(tokens, Token{
		Type: EOF, Lexeme: "", Literal: "", Line: line,
	})

	// We are at the beginning of the next lexeme.
	return NewScanner(source, tokens)
}
