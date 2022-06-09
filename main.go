package main

import (
	"fmt"
	"os"

	"github.com/andey-robins/interpreter/repl"
)

func main() {
	fmt.Println("Welcome to the Monkey programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
