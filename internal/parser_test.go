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
		expectedName    string
	}{
		{"type", "int"},
		{"type", "string"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedLiteral, decl.TokenLiteral())
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

	// Asserts.
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

func TestParseBlock(t *testing.T) {
	input := `
{}
{
	type int
	int a
}
`

	lexer := NewLexer(input)
	parser := NewParser(lexer)

	ast := parser.Parse()

	require := require.New(t)
	require.NotNil(ast)
	require.Len(ast.Declarations, 2)

	tests := []struct {
		expectedDeclarations []Declaration
	}{
		{},
		{
			expectedDeclarations: []Declaration{&Type{
				Token: Token{
					Type:    tokens.TYPE,
					Literal: "type",
				},
				Name: "int",
			}, &Var{
				Token: Token{
					Type:    tokens.IDENT,
					Literal: "int",
				},
				Type: "int",
				Name: "a",
			}},
		},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedDeclarations, decl.(*Block).Declarations, idx)
	}
}

func TestParseReturn(t *testing.T) {
	input := `
return a
return 1
`

	// Setup.
	lexer := NewLexer(input)
	parser := NewParser(lexer)

	// Test.
	ast := parser.Parse()

	// Asserts.
	require := require.New(t)
	require.NotNil(ast)

	tests := []struct {
		expectedReturnValue string
	}{
		{"a"},
		{"1"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedReturnValue, decl.(*Return).Value, idx)
	}
}
