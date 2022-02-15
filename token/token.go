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
	ASSIGN    = "="
	ADD       = "+"
	SUBSTRACT = "-"
	BANG      = "!"
	ASTERISK  = "*"
	BSLASH    = "/"
	LESSER    = "<"
	GREATER   = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	OPAREN    = "("
	CPAREN    = ")"
	OBRACE    = "{"
	CBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fun": FUNCTION,
	"let": LET,
}

// Lookup for identifier. Returns it's token if found one.
func LookupIdentifier(idn string) TokenType {
	if tk, ok := keywords[idn]; ok {
		return tk
	}

	return IDENT
}
