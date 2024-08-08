package main

import (
	"fmt"
	"go_interpreter/src/monkey/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey progamming language!\n", usr.Username)
	fmt.Printf("You can begin typing any commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
