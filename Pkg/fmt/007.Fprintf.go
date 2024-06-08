package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./007.example.txt")
	if err != nil {
		fmt.Print("Error in creating a file", err)
	}
	var name string = "Arvinder"
	var age int = 30
	_, err = fmt.Fprintf(file, "Hello! My name is %s. I am %d years old.", name, age)
	if err != nil {
		fmt.Print("Error in writing to a file.", err)
	}

	fmt.Print(`
	----------------- Writing to a buffer ----------------- 
	`)
	var buffer bytes.Buffer
	_, err = fmt.Fprintf(&buffer, "Hello! My name is %s. I am %d years old.", name, age)
	if err != nil {
		fmt.Print("Error in writing to a buffer", err)
	}
	fmt.Print(buffer.String())
}

/*
Fprintf is a function in the fmt package in Go that formats according to a format specifier
and writes to the specified io.Writer. This is useful for writing formatted output to files,
network connections, buffers, or other output streams that implement the io.Writer interface.


-> Fprintf formats its arguments based on a format specifier string.
-> The format specifier string contains verbs that define the format for each argument.
-> It writes the formatted output to the specified io.Writer.
*/
