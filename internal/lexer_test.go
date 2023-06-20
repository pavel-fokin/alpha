package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	input := `
type int

var num int

func add(int a, int b) int`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{tokens.TYPE, "type"},
		{tokens.IDENT, "int"},
		{tokens.VAR, "var"},
		{tokens.IDENT, "num"},
		{tokens.IDENT, "int"},
		{tokens.FUNC, "func"},
		{tokens.IDENT, "add"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "int"},
		{tokens.IDENT, "a"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "int"},
		{tokens.IDENT, "b"},
		{tokens.RPAREN, ")"},
		{tokens.IDENT, "int"},
		{tokens.EOF, ""},
	}

	lexer := NewLexer(input)

	for _, test := range tests {
		token := lexer.NextToken()
		require.Equal(t, test.expectedType, token.Type)
		require.Equal(t, test.expectedLiteral, token.Literal)
	}
}
