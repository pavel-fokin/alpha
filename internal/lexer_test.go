package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLexer_Type(t *testing.T) {
	input := `type int`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{tokens.TYPE, "type"},
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

func TestLexer_Var(t *testing.T) {
	input := `var num int`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{tokens.VAR, "var"},
		{tokens.IDENT, "num"},
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

func TestLexer_Func(t *testing.T) {
	input := `func add(a int, b int) int`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{tokens.FUNC, "func"},
		{tokens.IDENT, "add"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "a"},
		{tokens.IDENT, "int"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "b"},
		{tokens.IDENT, "int"},
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
