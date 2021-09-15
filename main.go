// Simple wordcount program.
// Part of this course -- https://stepik.org/course/96832/
// It ignores errors and counts all words in all strings
// that are passed as arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var counter int
	// Take only arguments and iterate over them.
	inputStrings := os.Args[1:]
	for _, arg := range inputStrings {
		// While iterating, split strings into words
		// using strings.Fields() and count nimber of fields.
		counter += len(strings.Fields(arg))
	}

	fmt.Print(counter)
}
