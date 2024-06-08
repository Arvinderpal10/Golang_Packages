package main

import (
	"fmt"
)

func main() {

	var name string = "Arvinder"
	var age int = 40
	var second string = "two"

	Output := fmt.Sprintln("Hello ! My name is ", name, ". I am ", age, "years old.")
	newOutput := fmt.Sprintln("This is line number ", second, " and the text is printed in newline.")

	fmt.Print(Output)
	fmt.Print(newOutput)

}

/*
-> Sprintln formats its arguments using the default formats.
-> It returns the formatted string with a newline character appended at the end.
-> Spaces are added between operands when neither is a string.
*/

/*
-> String Construction: Useful for building strings that will be used later.
-> Log Messages: Construct log messages with newline termination.
-> Returning Formatted Strings: When a function needs to return a formatted string that includes a newline.

*/
/*
Output :
Hello ! My name is  Arvinder . I am  40 years old.
This is line number  two  and the text is printed in newline.
*/
