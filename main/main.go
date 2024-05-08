package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"ascii_art"
)

// this function called ascii art
func main() {
	// Declaration
	const template = "../txt/standard.txt"

	// Case of multiple arguments
	if len(os.Args) != 2 {
		fmt.Println("Error, Please enter one argument")
		os.Exit(1)
	}
	arg := os.Args[1]

	// The given argument is an empty string.
	if len(arg) == 0 {
		return
	}

	fp, err := os.Open(template)
	if err != nil {
		log.Fatal(err)
	}
	// Close the file at the end.
	defer fp.Close()
	content, err := io.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	ascii_art.ArgProcessor(arg, string(content))
}
