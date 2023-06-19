package internal

import (
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
type int
type string

var num int
var str string`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{tokens.TYPE, "type"},
		{tokens.IDENT, "int"},
		{tokens.TYPE, "type"},
		{tokens.IDENT, "string"},
		{tokens.VAR, "var"},
		{tokens.IDENT, "num"},
		{tokens.IDENT, "int"},
		{tokens.VAR, "var"},
		{tokens.IDENT, "str"},
		{tokens.IDENT, "string"},
		{tokens.EOF, ""},
	}

	lexer := NewLexer(input)

	for idx, test := range tests {
		token := lexer.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				idx, test.expectedType, token.Type)
		}
		if token.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				idx, test.expectedLiteral, token.Literal)
		}
	}
}
