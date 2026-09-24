// challenges/types-primitive/begin/main.go
package main

import "fmt"

/* Task: see challenge.md, don't go by the comments here.
- Also, the sample output doesn't match the task for 4.
- 6. local val isn't a pointer, but the sample output shows a reference
- 7. is making me think that 2. was supposed to create a pointer, then work with that
- ...I made 'local val' a pointer (although the instructor hasn't demonstrated use of the "new" keyword yet)

Solution:
- local val was just a simple int
- and...he basically implemented what I had originally before I started analyzing the challenge.md
*/

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	var val *int = new(int)
	*val = 42

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", *val, *val)

	// print the value of the package-level variable "val"
	printGlobalVal()

	// update the package-level variable "val" and print it again
	updateGlobalVal()
	printGlobalVal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", val, val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*val = 100
	fmt.Printf("local val = %v\n", *val)
}

func printGlobalVal() {
	fmt.Printf("%T, global val = %v\n", val, val)
}

func updateGlobalVal() {
	val = "Updated global"
}
