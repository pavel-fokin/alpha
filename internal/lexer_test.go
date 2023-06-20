package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	for _, test := range tests {
		token := lexer.NextToken()
		require.Equal(t, test.expectedType, token.Type)
		require.Equal(t, test.expectedLiteral, token.Literal)
	}
}
