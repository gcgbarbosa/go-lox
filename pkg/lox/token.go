package lox

type TokenType int

const (
	LEFT_PAREN TokenType = iota // iota starts at 0 and increments for each constant
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMMA
	DOT
	MINUS
)

// more info on maligned structs:
// https://medium.com/@sebassegros/golang-dealing-with-maligned-structs-9b77bacf4b97
type Token struct {
	Literal interface{}
	Lexeme  string
	Line    int
	Type_   TokenType
}
