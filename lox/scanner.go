package lox

// import "fmt"

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

func (s *Scanner) Tokens() []Token {
	return s.tokens
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

	peek := func() rune {
		if isAtEnd() {
			return '\x00'
		}
		return rune(source[current])
	}

	// only consume the current character if it is what we are looking for
	match := func(expected rune) bool {
		if isAtEnd() {
			return false
		}

		if rune(source[current]) != expected {
			return false
		}

		current++
		return true
	}

	advance := func() rune {
		current += 1

		runes := []rune(source)

		// fmt.Println("Current:", current)

		return runes[current-1]
	}

	addTokenAndLiteral := func(type_ TokenType, literal interface{}) {
		text := string(source[start:current])
		tokens = append(tokens, Token{Type: type_, Lexeme: text, Literal: literal, Line: line})
	}

	// helper function that always passes nil for the literal
	addToken := func(type_ TokenType) {
		addTokenAndLiteral(type_, nil)
	}

	scanToken := func() {
		c := advance()

		switch c {

		// tokens
		case '(':
			addToken(LEFT_PAREN)
		case ')':
			addToken(RIGHT_PAREN)
		case '{':
			addToken(RIGHT_BRACE)
		case '}':
			addToken(LEFT_BRACE)
		case ',':
			addToken(COMMA)
		case '.':
			addToken(DOT)
		case '-':
			addToken(MINUS)
		case '+':
			addToken(PLUS)
		case ';':
			addToken(SEMICOLON)
		case '*':
			addToken(STAR)

		// operators
		// handles != or !
		case '!':
			if match('=') {
				addToken(BANG_EQUAL)
			} else {
				addToken(BANG)
			}

		// handles == or =
		case '=':
			if match('=') {
				addToken(EQUAL_EQUAL)
			} else {
				addToken(EQUAL)
			}

		// handles <= or =
		case '<':
			if match('=') {
				addToken(LESS_EQUAL)
			} else {
				addToken(LESS)
			}

		// handles >= or >
		case '>':
			if match('=') {
				addToken(GREATER_EQUAL)
			} else {
				addToken(GREATER)
			}

		// handles comments
		case '/':
			if match('/') {
				for peek() != '\n' && !isAtEnd() {
					advance()
				}
			} else {
				addToken(SLASH)
			}

		case ' ':
    case '\r':
    case '\t':

    case '\n':
      line += 1


		// Add more cases for other tokens like operators, literals, etc.
		default:
			// Handle unknown characters or skip whitespace here
			error(line, "Unexpected character.")
			// fmt.Println("Unknown character:", c)
		}
	}

	for !isAtEnd() { // while !isAtEnd()
		// if isAtEnd() {
		// 	break
		// }

		scanToken()

		start = current
	}

	// add EOF to end read
	tokens = append(tokens, Token{
		Type: EOF, Lexeme: "", Literal: "", Line: line,
	})

	return NewScanner(source, tokens)
}
