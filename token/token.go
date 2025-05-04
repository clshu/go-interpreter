package token

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
	// BANG represents a bang token.
	BANG Type = "!"
	// ASTERISK represents an asterisk token.
	ASTERISK Type = "*"
	// SLASH represents a slash token.
	SLASH Type = "/"
	// LT represents a less than token.
	LT Type = "<"
	// GT represents a greater than token.
	GT Type = ">"
	// EQ represents an equal token.
	EQ Type = "=="
	// NotEQ represents a not equal token.
	NotEQ Type = "!="

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
	// TRUE represents a true keyword token.
	TRUE Type = "TRUE"
	// FALSE represents a false keyword token.
	FALSE Type = "FALSE"
	// IF represents an if keyword token.
	IF Type = "IF"
	// ELSE represents an else keyword token.
	ELSE Type = "ELSE"
	// RETURN represents a return keyword token.
	RETURN Type = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Type represents the type of token.
type Type string

// Token defines a token with its type and literal value.
type Token struct {
	Type    Type
	Literal string
}

// LookupIdent checks if the identifier is a keyword and returns its type.
// If not, it returns IDENT.
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
