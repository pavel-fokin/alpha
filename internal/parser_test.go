package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	input := `
type int
type string

var num int
var str string
`

	lexer := NewLexer(input)
	parser := NewParser(lexer)

	ast := parser.Parse()
	if ast == nil {
		t.Fatal("Parse() returned 'nil'.")
	}
	if len(ast.Declarations) != 4 {
		t.Fatal("ast.Declarations doesn't have 4 declarations.")
	}

	tests := []struct {
		expectedLiteral string
		expectedType    Declaration
		expectedName    string
	}{
		{"type", &Type{}, "int"},
		{"type", &Type{}, "string"},
		{"var", &Var{}, "str"},
		{"var", &Var{}, "num"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(t, test.expectedLiteral, decl.TokenLiteral())
		require.IsType(t, test.expectedType, decl)
		// require.Equal(t, test.expectedName, decl.(test.expectedType).Name)
	}
}

func TestParseFunc(t *testing.T) {
	input := `
func foo()
func foo(a)
func foo() int
func foo(a, b) int
`

	lexer := NewLexer(input)
	parser := NewParser(lexer)

	ast := parser.Parse()

	require := require.New(t)
	require.NotNil(ast)

	tests := []struct {
		expectedLiteral string
		expectedName    string
		expectedParams  []string
		expectedReturns []string
	}{
		{"func", "foo", []string{}, []string{}},
		{"func", "foo", []string{"a"}, []string{}},
		{"func", "foo", []string{}, []string{"int"}},
		{"func", "foo", []string{"a", "b"}, []string{"int"}},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		require.Equal(test.expectedLiteral, decl.TokenLiteral(), idx)
		require.Equal(test.expectedName, decl.(*Func).Name, idx)
		require.Equal(test.expectedParams, decl.(*Func).Params, idx)
		require.Equal(test.expectedReturns, decl.(*Func).Returns, idx)
	}
}
