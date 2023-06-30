package internal

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tt TokenType, l string) Token {
	return Token{tt, l}
}

var tokens = struct {
	ILLEGAL TokenType
	EOF     TokenType

	IDENT TokenType

	// Brackets.
	COMMA  TokenType
	LPAREN TokenType
	RPAREN TokenType

	// Keywords.
	TYPE TokenType
	VAR  TokenType
	FUNC TokenType
}{
	"ILLEGAL",
	"EOF",
	"IDENT",
	",",
	"(",
	")",
	"TYPE",
	"VAR",
	"FUNC",
}

var TokenKeywords = map[string]TokenType{
	"type": tokens.TYPE,
	"var":  tokens.VAR,
	"func": tokens.FUNC,
}

func LookupKeyword(ident string) TokenType {
	if tokenType, ok := TokenKeywords[ident]; ok {
		return tokenType
	}
	return tokens.IDENT
}
