// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
)

/*
Task: see challenge.md
Solution:
* library: instructor used a map[string][]book (I was wondering where maps came into this exercise).
	-	I would've called it 'bibliography' or 'collection'
	- We also haven't learned the 'range' keyword yet, so I guess that's cheating, haha
	- He created a single 'author' object to create the two books (which is a good practice)
	- Apparently the function 'spew' was pre-existing in the module file?  Why would we look there?
*/

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) ([]book, error) {
	list, ok := l[name]

	if !ok {
		return nil, fmt.Errorf("No books by %v in the library!", name)
	}
	return list, nil
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	lib.addBook(
		book{
			title:  "The Will of the Many",
			author: author{name: "James Islington"},
		})

	lib.addBook(
		book{
			title:  "The Strength of the Few",
			author: author{name: "James Islington"},
		})
	// add 1 book to the library by a different author
	lib.addBook(
		book{
			title:  "The Way of Kings",
			author: author{name: "Brandon Sanderson"},
		})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	books, err := lib.lookupByAuthor("James Islington")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// print out the first book's title and its author's name
	fmt.Printf("First book: %#v\n", books[0])
}
