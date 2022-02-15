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
	IS_EQUAL  = "=="
	NOT_EQUAL = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Lookup for identifier. Returns it's token if found one.
func LookupIdentifier(idn string) TokenType {
	if tk, ok := keywords[idn]; ok {
		return tk
	}

	return IDENT
}
