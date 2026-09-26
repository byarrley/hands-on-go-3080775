// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

/*
Task: see challenge.md
Notes:
- I ended up using regex instead of a switch statement to check the type of each character for the initial implementation
*/

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered: %v", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fnm := os.Args
	fmt.Printf("args %#v...\n", fnm)

	if len(fnm) < 2 {
		panic("Must provide filename to count!\n")
	}
	fmt.Printf("Counting %s...\n", fnm[1])

	// convert the bytes to a string
	bytes, err := os.ReadFile(fnm[1])

	if err != nil {
		panic(err)
	}
	contents := string(bytes)

	// initialize a map to store the counts
	counts := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(contents)

	// capture the length of the words slice
	counts["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	counts["letters"] = regexCount(`[[:alpha:]]`, contents)
	counts["numbers"] = regexCount(`[[:digit:]]`, contents)
	counts["symbols"] = regexCount("[!-/:-@[-`{-~]", contents)
	// dump the map to the console using the spew package
	spew.Dump(counts)
}

func regexCount(exp string, s string) int {
	result := 0
	re := regexp.MustCompile(exp)

	matches := re.FindAllString(s, -1)
	fmt.Printf("\nre=%#v\n", re)
	fmt.Printf("s=%#v\n", s)
	fmt.Printf("matches=%q\n", matches)
	for _, m := range matches {
		result += len(m)
	}
	return result
}
