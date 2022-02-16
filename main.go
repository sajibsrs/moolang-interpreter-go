package main

import (
	"fmt"
	"moolang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to moolang REPL!\n", user.Username)
	fmt.Print("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
