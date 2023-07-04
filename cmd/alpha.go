package main

import (
	"os"
	"os/user"

	"github.com/pavel-fokin/alpha/internal/cli"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		cli.Run(os.Args[1])
		return
	}

	cli.NewREPL(user.Username).Start(os.Stdin, os.Stdout)
}
