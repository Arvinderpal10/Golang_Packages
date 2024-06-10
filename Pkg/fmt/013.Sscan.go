package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var country string

	// Correct input format
	input := "Arvinder 25 India" //

	// Incorrect input format (missing country)
	incorrectInput1 := "Arvinder 28"           // Error : EOF
	incorrectInput2 := "Arvinder 28  "         // Error : EOF
	incorrectInput3 := "Arvinder Pal 28 34.56" //Error: expected integer
	incorrectInput4 := "100 28  30.40"         // Working

	// Function to scan input and handle errors
	scanInput := func(input string) {
		n, err := fmt.Sscan(input, &name, &age, &country)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Total %d items. Name: %s, Age: %d, Country: %s\n", n, name, age, country)
		}
	}

	// Scanning the correct input
	fmt.Println("Scanning correct input:")
	scanInput(input)

	// Scanning the incorrect input
	fmt.Println("\nScanning incorrect input:")
	scanInput(incorrectInput1)
	scanInput(incorrectInput2)
	scanInput(incorrectInput3)
	scanInput(incorrectInput4)
}

/*
OUTPUT :
Scanning correct input:
Total 3 items. Name: Arvinder, Age: 25, Country: India

Scanning incorrect input:
Error: EOF
Error: EOF
Error: expected integer
Total 3 items. Name: 100, Age: 28, Country: 30.40

--------//----------//-----------
Sscan in Go is a function from the fmt package used to scan space-separated values from a string
and store them into successive arguments. It's particularly useful for parsing strings containing
structured data.

-> fmt.Sscan reads successive space-separated values from the provided string and stores
them in successive arguments.
-> It returns the number of items successfully scanned and an error, if any.


====> Common Use Cases
-> Parsing Command-Line Arguments: Reading and parsing arguments provided in a command-line tool.
-> Processing Data from Files: Scanning structured data read from files.
_> Unit Tests: Simulating user input or other string-based data for testing purposes.
*/
