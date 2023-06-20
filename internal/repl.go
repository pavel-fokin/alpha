package internal

import (
	"bufio"
	"fmt"
	"io"
)

const PROMT = ">> "

type REPL struct{}

func (r REPL) Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := NewLexer(line)
		parser := NewParser(lexer)
		ast := parser.Parse()

		io.WriteString(out, ast.String())
		io.WriteString(out, "\n")
	}
}
