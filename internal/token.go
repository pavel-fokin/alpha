package internal

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var tokens = struct {
	ILLEGAL TokenType
	EOF     TokenType
	IDENT   TokenType
	TYPE    TokenType
}{
	"ILLEGAL",
	"EOF",
	"IDENT",
	"TYPE",
}

var TokenKeywords = map[string]TokenType{
	"type": tokens.TYPE,
}

func LookupKeyword(ident string) TokenType {
	if tokenType, ok := TokenKeywords[ident]; ok {
		return tokenType
	}
	return tokens.IDENT
}
