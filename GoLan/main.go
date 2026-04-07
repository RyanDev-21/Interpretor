package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/RyanDev-21/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello welcome to the monkey language!\n %s", user.Name)
	fmt.Println("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
