package main

import (
	"bytes" // writing to buffer
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("./006.example.txt")
	if err != nil {
		fmt.Print("Error in creating new file.", err)
		return
	}
	//Closing file before the function exits
	defer file.Close()
	var name string = "Arvinder"
	var age int = 30
	_, err = fmt.Fprint(file, "Hello! My name is ", name, ". I am ", age, " years old")
	if err != nil {
		fmt.Print("Error writin to file.", err)
	}

	/*
		Output : Hello! My name is Arvinder. I am 30 years old --> Will be written to file (example.txt)
	*/

	fmt.Print(`
	----------- Writing to a Buffer ---------------
	`)

	var buffer bytes.Buffer
	var myName string = "Arvinder"
	var myAge int = 30
	_, err = fmt.Fprint(&buffer, "Hello! My name is ", myName, ". I am ", myAge, " years Old.")
	if err != nil {
		fmt.Print("Error in writing  to buffer", err)
	}

	fmt.Print(buffer.String())

	// Output :
	//----------- Writing to a Buffer ---------------
	// Hello! My name is Arvinder. I am 30 years Old.

}

/*
Fprint is a function in the fmt package in Go that formats (based on default format) its arguments using the default
formats and writes the result to a specified io.Writer. This function is useful for writing
formatted output to files, network connections, or other output streams.


-> Fprint formats its arguments using the default formats.
-> It writes the formatted output to the specified io.Writer.
-> Spaces are added between operands when neither is a string.

*/

// Common functions related to Buffer
// Write strings to the buffer
// buffer.WriteString("Hello!")
// buffer.WriteString(" You are writing to a buffer.")

// // reading the content from buffer

// fmt.Print(buffer.String())

// // Updating the content of buffer
// buffer.Write([]byte(" You are an awesome learner."))
// fmt.Print(buffer.String())
