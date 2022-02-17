package parser

import (
	"moolang/ast"
	"moolang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	lex := lexer.New(input)
	prc := New(lex)

	program := prc.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does't contain 3 statements. Got=%d", len(program.Statements))
	}

	tests := []struct {
		exprectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, ts := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, ts.exprectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. Got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. Got=%s", name, letStmt.Name)
		return false
	}

	return true
}
