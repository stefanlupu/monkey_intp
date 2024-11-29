package main

import (
	"fmt"
	"monkeyintp/repl"
	"os"
)

func main() {
	fmt.Println("* Monkey REPL:")
	repl.Start(os.Stdin, os.Stdout)
}
