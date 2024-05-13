package main

import (
	"fmt"
	"os"

	"ascii_art"
)

func main() {
	// Constents Declaration.
	const templatePath = "../txt/standard.txt"

	// Case of multiple arguments
	if len(os.Args) != 2 {
		fmt.Println("Error, Please enter one argument")
		os.Exit(1)
	}
	arg := os.Args[1]

	// The given argument is an empty string
	if len(arg) == 0 {
		return
	}

	fmt.Print(ascii_art.ArgProcessor(arg, templatePath))
}
