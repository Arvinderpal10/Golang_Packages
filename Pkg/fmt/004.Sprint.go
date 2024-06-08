package main

import (
	"fmt"
)

func main() {
	var name string = "Arvinder"
	var age int = 30
	output := fmt.Sprint("Hello! My name is ", name, ". I am ", age, " years old")

	fmt.Println(output)

	fmt.Println("--------------------------------")
	/* newline is not appended at the end */
	var first string = "One"
	var second string = "Two"
	first_line := fmt.Sprint("This is First Number : ", first)
	second_line := fmt.Sprint("This is line number :", second)
	fmt.Print(first_line)
	fmt.Print(second_line)

	// Output : This is First Number : OneThis is line number :Two
	// Will not append NewLine

}

/*
-> Sprint formats its arguments using the default formats.
-> It returns the formatted string without printing it.
-> Spaces are added between operands when neither is a string.

Common Use Cases:
1. String Concatenation: Sprint can be used to concatenate multiple values into a single string with formatting.
2. Building Log Messages: Useful for constructing log messages that include various data types.
3. Returning Formatted Strings: When you need to return a formatted string from a function rather than printing it immediately.
*/

/*
OUTPUT :
Hello! My name is Arvinder. I am 30 years old

*/
