package lox

type TokenType int

// more info on maligned structs:
// https://medium.com/@sebassegros/golang-dealing-with-maligned-structs-9b77bacf4b97
type Token struct {
	Literal interface{}
	Lexeme  string
	Line    int
	Type    TokenType
}

const (
	// single-character tokens
	LEFT_PAREN TokenType = iota // iota starts at 0 and increments for each constant
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMMA
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
