package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/pavel-fokin/alpha/internal"
)

func Run(filename string) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("couldn't read file: %w", err)
	}

	lexer := internal.NewLexer(string(data))
	parser := internal.NewParser(lexer)
	ast := parser.Parse()

	io.WriteString(os.Stdout, ast.String())
	io.WriteString(os.Stdout, "\n")

	return nil
}
