package token

// Type represents the type of token.
type Type string

// Token defines a token with its type and literal value.
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL represents an illegal token.
	ILLEGAL Type = "ILLEGAL"
	// EOF represents the end of file token.
	EOF Type = "EOF"
	// IDENT represents an identifier token.
	IDENT Type = "IDENT" // add, foobar, x, y, ...
	// INT represents an integer token.
	INT Type = "INT" // 1234567890

	// operators

	// ASSIGN represents an assignment token.
	ASSIGN Type = "="
	// PLUS represents a plus token.
	PLUS Type = "+"
	// MINUS represents a minus token.
	MINUS Type = "-"

	// Delimiters

	// COMMA represents a comma token.
	COMMA Type = ","
	// SEMICOLON represents a semicolon token.
	SEMICOLON Type = ";"
	// LPAREN represents a left parenthesis token.
	LPAREN Type = "("
	// RPAREN represents a right parenthesis token.
	RPAREN Type = ")"
	// LBRACE represents a left brace token.
	LBRACE Type = "{"
	// RBRACE represents a right brace token.
	RBRACE Type = "}"

	// Keywords

	// FUNCTION represents a function keyword token.
	FUNCTION Type = "FUNCTION"
	// LET represents a let keyword token.
	LET Type = "LET"
)
