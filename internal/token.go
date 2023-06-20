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
	// Keywords.
	TYPE TokenType
	VAR  TokenType
	FUNC TokenType
}{
	"ILLEGAL",
	"EOF",
	"IDENT",
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
