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

		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 9;
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
		{token.BANG, "!"},
		{token.SUBSTRACT, "-"},
		{token.BSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESSER, "<"},
		{token.INT, "10"},
		{token.GREATER, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.OPAREN, "("},
		{token.INT, "5"},
		{token.LESSER, "<"},
		{token.INT, "10"},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.ELSE, "else"},
		{token.OBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.INT, "10"},
		{token.IS_EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	for i, ts := range tests {
		tk := lex.NextToken()

		if tk.Type != ts.exptedType {
			t.Logf("%d , %d", lex.position, lex.readPosition)
			t.Fatalf("tests[%d] - Wrong token type. Expected=%q, got=%q", i, ts.exptedType, tk.Type)
		}

		if tk.Literal != ts.expectedLiteral {
			t.Fatalf("tests[%d] - Wrong literal. Expected=%q, got=%q", i, ts.expectedLiteral, tk.Literal)
		}
	}
}
