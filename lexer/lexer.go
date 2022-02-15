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

// Reads current char from Lexer input.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Reads next char from Lexer input.
func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.O_PAREN, l.ch)
	case ')':
		tk = newToken(token.C_PAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.ADD, l.ch)
	case '{':
		tk = newToken(token.O_BRACE, l.ch)
	case '}':
		tk = newToken(token.C_BRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	}

	l.readChar()
	return tk
}

// Creates new token and returns Token object.
func newToken(tk token.TokenType, ch byte) token.Token {
	return token.Token{Type: tk, Literal: string(ch)}
}
