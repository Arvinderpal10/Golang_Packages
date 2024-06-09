package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./008.example.txt")
	if err != nil {
		fmt.Print("Error in creating a file", err)
	}
	var name string = "Arvinder"
	var age int = 30
	_, err = fmt.Fprintln(file, "Hello! My name is", name, ". I am", age, "years old.")
	if err != nil {
		fmt.Print("Error in writing to a file.", err)
	}

	fmt.Print(`
	----------------- Writing to a buffer ----------------- 
	`)
	var buffer bytes.Buffer
	_, err = fmt.Fprintln(&buffer, "Hello! My name is", name, ". I am", age, "years old.")
	if err != nil {
		fmt.Print("Error in writing to a buffer", err)
	}
	fmt.Print(buffer.String())
}

/*
Fprintln is a function in the fmt package in Go that formats its arguments using the default
formats and writes the result to the specified io.Writer, followed by a newline. This is useful
for writing formatted output to files, network connections, buffers, or other output streams
that implement the io.Writer interface, with the convenience of automatically appending a
newline at the end.


-> Fprintln formats its arguments using the default formats.
-> It writes the formatted output to the specified io.Writer.
-> A newline character is appended at the end of the output.
-> Spaces are added between operands when neither is a string.

*/
