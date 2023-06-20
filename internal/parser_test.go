package internal

import (
	"testing"
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
		expectedTypeName string
	}{
		{"int"},
		{"string"},
	}

	for idx, test := range tests {
		decl := ast.Declarations[idx]
		if decl.TokenLiteral() != "type" {
			t.Fatalf("Type declaration TokenLiteral not 'type' got='%s'", test.expectedTypeName)
		}
		typeDecl, ok := decl.(*Type)
		if !ok {
			t.Fatalf("not a *Type, got=%T", decl)
		}

		if typeDecl.Name != test.expectedTypeName {
			t.Fatalf("type.Name not '%s', got='%s'", test.expectedTypeName, typeDecl.Name)
		}
	}
}
