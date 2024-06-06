package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello")
	fmt.Print("World.")
}

/*
- Print formats its arguments using the default formats and writes the output to standard
output (usually the console).
- It does not add any spaces between the arguments unless one of the arguments is a string
that includes spaces.
- It does not automatically add a newline at the end of the output, so subsequent output
will be on the same line.
*/

/*
Output : HelloWorld
*/
