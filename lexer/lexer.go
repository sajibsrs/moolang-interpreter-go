package lexer

import (
	"moolang/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}

// Creates new Lexer object and returns a pointer to it.
func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

// Reads current char from Lexer input and moves to next position.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Reads next char from Lexer input and turns current Token.
func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.OPAREN, l.ch)
	case ')':
		tk = newToken(token.CPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.ADD, l.ch)
	case '-':
		tk = newToken(token.SUBSTRACT, l.ch)
	case '*':
		tk = newToken(token.ASTERISK, l.ch)
	case '/':
		tk = newToken(token.BSLASH, l.ch)
	case '<':
		tk = newToken(token.LESSER, l.ch)
	case '>':
		tk = newToken(token.GREATER, l.ch)
	case '!':
		tk = newToken(token.BANG, l.ch)
	case '{':
		tk = newToken(token.OBRACE, l.ch)
	case '}':
		tk = newToken(token.CBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdentifier(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Type = token.INT
			tk.Literal = l.readNumber()
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tk
}

// Reads a number and returns it.
// Then moves to the next character.
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Reads an identifier and returns it.
// Then moves to the next character.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Ignore whitespace characters (Space, TAB and Carriage return)
// and move to next character.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

// Creates new token and returns Token object.
func newToken(tk token.TokenType, ch byte) token.Token {
	return token.Token{Type: tk, Literal: string(ch)}
}

// Checks if passed argument is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Checks if passed argument is a numerical digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
