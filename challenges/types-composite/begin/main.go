// challenges/types-composite/begin/main.go
package main

import "fmt"

/*
Task: see challenge.md
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
type library struct {
	books []book
}

// define addBook to add a book to the library
func (l *library) addBook(b book) {
	l.books = append(l.books, b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	result := make([]book, 0)

	for _, b := range l.books {
		if b.author.name == name {
			result = append(result, b)
		}
	}
	return result
}

func (l library) spew() {
	fmt.Println("spewing...")
	for _, b := range l.books {
		fmt.Printf("%#v\n", b)
	}
	fmt.Println("done")
}

func main() {
	// create a new library
	var lib library

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
	lib.spew()

	// lookup books by known author in the library
	books := lib.lookupByAuthor("James Islington")

	// print out the first book's title and its author's name
	fmt.Printf("First book: %#v\n", books[0])
}
