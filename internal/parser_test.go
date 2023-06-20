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

	program := parser.ParseProgram()
	if program == nil {
		t.Fatal("ParserProgram() returned 'nil'.")
	}
	if len(program.Declarations) != 4 {
		t.Fatal("program.Declarations doesn't have 2 declarations.")
	}

	tests := []struct {
		expectedTypeName string
	}{
		{"int"},
		{"string"},
	}

	for idx, test := range tests {
		decl := program.Declarations[idx]
		if decl.TokenLiteral() != "type" {
			t.Fatalf("Type declaration TokenLiteral not 'type' got='%s'", test.expectedTypeName)
		}
		typeDecl, ok := decl.(*TypeDeclaration)
		if !ok {
			t.Fatalf("not a *TypeDeclaration, got=%T", decl)
		}

		if typeDecl.Name != test.expectedTypeName {
			t.Fatalf("typeDecl.Name not '%s', got='%s'", test.expectedTypeName, typeDecl.Name)
		}
	}
}
