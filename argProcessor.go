package ascii_art

import (
	"fmt"
	"strings"
)

func ArgProcessor(arg, content string) {
	const artCharHeight = 8

	output := ""
	argWords := strings.Split(arg, "\\n")
	count := 0
	// Count the empty strings (\n\n)
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
		return
	}

	for i := 0; i < len(argWords); i++ {
		if argWords[i] == "" {
			output += "\n"
			continue
		}
		for line := 0; line < artCharHeight && argWords[i] != ""; line++ {
			output += Gathering(argWords[i], content, line)
		}
	}
	fmt.Print(output)
}
