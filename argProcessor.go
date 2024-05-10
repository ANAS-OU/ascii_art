package ascii_art

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/**
 * A function that gather the equivilant
 * of the input from the file
 *
 * Return : string
 */

func Gathering(input, content string, line int) string {
	//testing this commit
	standard := strings.Split(content, "\n")[1:]
	var slc []string
	for _, c := range input {
		if c >= ' ' && c <= '~' {
			j := int(c-' ')*9 + line
			slc = append(slc, standard[j])
		} else {
			fmt.Println("Error : input is invalid")
			os.Exit(1)
		}
	}
	result := strings.Join(slc, "")
	return result + "\n"
}

func ArgProcessor(arg string, templatePath string) string {
	// Declaration
	// const templatePath = "/txt/standard.txt"
	const artCharHeight = 8

	fp, err := os.Open(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file at the end.
	defer fp.Close()
	content, err := io.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	output := ""
	argWords := strings.Split(arg, "\\n")
	count := 0

	// Count the empty strings
	for i := 0; i < len(argWords); i++ {
		if argWords[i] == "" {
			count++
		} else {
			count = 0
			break
		}
	}

	if count >= 1 {
		for i := 0; i < count-1; i++ {
			fmt.Println()
		}
		return ""
	}

	for i := 0; i < len(argWords); i++ {
		if argWords[i] == "" {
			output += "\n"
			continue
		}
		for line := 0; line < artCharHeight && argWords[i] != ""; line++ {
			output += Gathering(argWords[i], string(content), line)
		}
	}
	return output
}
