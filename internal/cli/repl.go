package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pavel-fokin/alpha/internal"
)

const PROMT = ">> "

type REPL struct {
	username string
}

func NewREPL(username string) *REPL {
	return &REPL{
		username: username,
	}
}

func (r REPL) Start(in io.Reader, out io.Writer) {
	fmt.Printf("Hello %s! This is the Alpha(α) programming language!\n", r.username)

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := internal.NewLexer(line)
		parser := internal.NewParser(lexer)
		ast := parser.Parse()

		io.WriteString(out, ast.String())
		io.WriteString(out, "\n")
	}
}
