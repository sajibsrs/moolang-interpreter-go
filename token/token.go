package token

type TokenType string

// moolang token
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // Identifiers
	INT   = "INT"   // Integer

	// Operators
	ASSIGN = "="
	ADD    = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	O_PAREN   = "("
	C_PAREN   = ")"
	O_BRACE   = "{"
	C_BRACE   = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
