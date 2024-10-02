package lox

import "fmt"

type TokenType int

// more info on maligned structs:
// https://medium.com/@sebassegros/golang-dealing-with-maligned-structs-9b77bacf4b97
type Token struct {
	Literal interface{}
	Lexeme  string
	Line    int
	Type    TokenType
}

func ToString(t Token) string {
	var literalStr string

	// Check if t.Literal is a string or something else
	if t.Literal == nil {
		literalStr = "nil"
	} else if literal, ok := t.Literal.(string); ok {
		literalStr = literal
	} else {
		literalStr = fmt.Sprintf("%v", t.Literal)
	}

	return t.Type.String() + " " + t.Lexeme + " " + literalStr
}

const (
	// single-character tokens
	LEFT_PAREN TokenType = iota // iota starts at 0 and increments for each constant
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (tt TokenType) String() string {
	switch tt {
	// Single-character tokens
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case SLASH:
		return "/"
	case STAR:
		return "*"

	// One or two character tokens
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="

	// Literals
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"

	// Keywords
	case AND:
		return "and"
	case CLASS:
		return "class"
	case ELSE:
		return "else"
	case FALSE:
		return "false"
	case FUN:
		return "fun"
	case FOR:
		return "for"
	case IF:
		return "if"
	case NIL:
		return "nil"
	case OR:
		return "or"
	case PRINT:
		return "print"
	case RETURN:
		return "return"
	case SUPER:
		return "super"
	case THIS:
		return "this"
	case TRUE:
		return "true"
	case VAR:
		return "var"
	case WHILE:
		return "while"

	// End of file
	case EOF:
		return "EOF"

	default:
		return "UNKNOWN"
	}
}
