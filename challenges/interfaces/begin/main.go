// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

/*
Task: See challenge.md
Notes:
  - I think we can reuse the unicode filters from the previous challenge for the implementation of "count"
  - The expected output for this challenge doesn't match the expected output for the previous challenge, despite what the text says
  - Assuming that the input data is still challenges/data.txt, I think there is an issue with the expected output for this challenge.
  - 'wc ../../data.txt' returns '0  427 2296 ../../data.txt'; if we add a type for counting whitespace, the total number of bytes shouldn't exceed 2296
    (which the expected output does, there's some double counting or some other error)

Output:

	(map[string]int) (len=5) {
	 (string) (len=5) "words": (int) 427,
	 (string) (len=7) "letters": (int) 1784,
	 (string) (len=7) "numbers": (int) 2,
	 (string) (len=7) "symbols": (int) 84,
	 (string) (len=6) "spaces": (int) 426
	}

Which gives us 2296 bytes.

Solution:
- Instructor made the same mistake I made by counting "symbols" as "!letters" and "!numbers", which would count spaces
(since we're operating on the full string, and not individual words).  Note: "symbols" + "spaces" = 510, which is the same as the expected number of "symbols"
*/
type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }

func (l letterCounter) name() string {
	return l.identifier
}
func (l letterCounter) count(input string) int {
	result := 0
	for _, c := range input {
		if unicode.IsLetter(c) {
			result++
		}
	}
	return result
}

type numberCounter struct{ designation string }

func (n numberCounter) name() string {
	return n.designation
}
func (n numberCounter) count(input string) int {
	result := 0
	for _, c := range input {
		if unicode.IsNumber(c) {
			result++
		}
	}
	return result
}

type symbolCounter struct{ label string }

func (s symbolCounter) name() string {
	return s.label
}
func (s symbolCounter) count(input string) int {
	result := 0
	for _, c := range input {
		if unicode.IsSymbol(c) || unicode.IsPunct(c) {
			result++
		}
	}
	return result
}

type whitespaceCounter struct{ etiqueta string }

func (w whitespaceCounter) name() string {
	return w.etiqueta
}
func (w whitespaceCounter) count(input string) int {
	result := 0
	for _, c := range input {
		if unicode.IsSpace(c) {
			result++
		}
	}
	return result
}

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	// spew.Dump(data)

	// call doAnalysis and pass in the data and the counters
	analysis := doAnalysis(data,
		letterCounter{identifier: "letters"},
		numberCounter{designation: "numbers"},
		symbolCounter{label: "symbols"},
		whitespaceCounter{etiqueta: "spaces"},
	)
	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
