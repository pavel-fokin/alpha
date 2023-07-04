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
	LBRACE TokenType
	RBRACE TokenType

	// Keywords.
	TYPE   TokenType
	FUNC   TokenType
	RETURN TokenType
}{
	"ILLEGAL",
	"EOF",
	"IDENT",
	",",
	"(",
	")",
	"{",
	"}",
	"TYPE",
	"FUNC",
	"RETURN",
}

var TokenKeywords = map[string]TokenType{
	"type":   tokens.TYPE,
	"func":   tokens.FUNC,
	"return": tokens.RETURN,
}

func LookupKeyword(ident string) TokenType {
	if tokenType, ok := TokenKeywords[ident]; ok {
		return tokenType
	}
	return tokens.IDENT
}
