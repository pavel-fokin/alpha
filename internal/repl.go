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
		if line == "exit()" {
			return
		}
		lexer := NewLexer(line)
		parser := NewParser(lexer)
		program := parser.ParseProgram()

		for tok := lexer.NextToken(); tok.Type != tokens.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}

}
