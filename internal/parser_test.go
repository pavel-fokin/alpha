package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseType(t *testing.T) {
	input := `
type int
type string`

	lexer := NewLexer(input)
	parser := NewParser(lexer)

	ast := parser.Parse()

	require := require.New(t)
	require.NotNil(ast)

	tests := []struct {
		expectedLiteral string
		expectedType    Declaration
		expectedName    string
	}{
		{"type", &Type{}, "int"},
		{"type", &Type{}, "string"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedLiteral, decl.TokenLiteral())
		require.IsType(test.expectedType, decl)
		require.Equal(test.expectedName, decl.(*Type).Name)
	}
}

func TestParseVar(t *testing.T) {
	input := `
int num
string str`

	// Setup.
	lexer := NewLexer(input)
	parser := NewParser(lexer)

	// Test.
	ast := parser.Parse()

	// Assert.
	require := require.New(t)
	require.NotNil(ast)

	tests := []struct {
		expectedType string
		expectedName string
	}{
		{"int", "num"},
		{"string", "str"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedType, decl.(*Var).Type, idx)
		require.Equal(test.expectedName, decl.(*Var).Name, idx)
	}
}

func TestParseFunc(t *testing.T) {
	input := `
func foo()
func foo(int a)
func foo() int
func foo(int a, int b) int
`

	lexer := NewLexer(input)
	parser := NewParser(lexer)

	ast := parser.Parse()

	require := require.New(t)
	require.NotNil(ast)

	tests := []struct {
		expectedLiteral string
		expectedName    string
		expectedArgs    []*Var
		expectedReturns []string
	}{
		{"func", "foo", []*Var{}, []string{}},
		{"func", "foo", []*Var{
			{
				Token: Token{
					Type:    tokens.IDENT,
					Literal: "int",
				},
				Name: "a",
				Type: "int",
			},
		}, []string{}},
		{"func", "foo", []*Var{}, []string{"int"}},
		{"func", "foo", []*Var{
			{
				Token: Token{
					Type:    tokens.IDENT,
					Literal: "int",
				},
				Name: "a",
				Type: "int",
			},
			{
				Token: Token{
					Type:    tokens.IDENT,
					Literal: "int",
				},
				Name: "b",
				Type: "int",
			},
		}, []string{"int"}},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedLiteral, decl.TokenLiteral(), idx)
		require.Equal(test.expectedName, decl.(*Func).Name, idx)
		require.Equal(test.expectedArgs, decl.(*Func).Args, idx)
		require.Equal(test.expectedReturns, decl.(*Func).Returns, idx)
	}
}
