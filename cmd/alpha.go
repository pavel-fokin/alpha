package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/pavel-fokin/alpha/internal"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Alpha(α) programming language!\n", user.Username)
	internal.REPL{}.Start(os.Stdin, os.Stdout)
}
