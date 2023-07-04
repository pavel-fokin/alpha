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

	cli.NewREPL(user.Username).Start(os.Stdin, os.Stdout)
}
