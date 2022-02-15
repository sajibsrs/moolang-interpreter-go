package lexer

import (
	"moolang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		exptedType      token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.O_PAREN, "("},
		{token.C_PAREN, ")"},
		{token.O_BRACE, "{"},
		{token.C_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	for i, ts := range tests {
		tk := lex.NextToken()

		if tk.Type != ts.exptedType {
			t.Fatalf("tests[%d] - Wrong token type. Expected=%q, got=%q", i, ts.exptedType, tk.Type)
		}

		if tk.Literal != ts.expectedLiteral {
			t.Fatalf("tests[%d] - Wrong literal. Expected=%q, got=%q", i, ts.expectedLiteral, tk.Literal)
		}
	}
}
