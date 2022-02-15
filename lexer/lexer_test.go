package lexer

import (
	"moolang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		
		let add = fun(x, y) {
			x + y;
		};

		let result = add(five, ten);
		`

	tests := []struct {
		exptedType      token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.OPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.OPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.CPAREN, ")"},
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
