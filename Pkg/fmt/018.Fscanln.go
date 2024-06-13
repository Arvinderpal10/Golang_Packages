package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("018.Fscanln.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var name string
	var age int
	n, err := fmt.Fscanln(file, &name, &age)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	fmt.Printf("Successfully scanned %d items\n", n)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

/*
-> Fscanln is a function used for formatted input that reads from a specified io.Reader
and stops scanning at a newline character. It's similar to Fscan but with a crucial difference
in how it handles newlines.

Purpose:
-> Reads formatted data from a source, parses it according to format specifiers, stores
extracted values in provided variables, and stops scanning at a newline character.

n, err := fmt.Fscanln(r, a...)

-> r (mandatory): An io.Reader interface representing the source of the formatted text. This
is usually a file opened using os.Open or similar functions.
-> a... (variadic): A sequence of variables where the extracted values will be stored. The
number and types of variables must match the format specifiers in the format string used when
parsing the input.

-> n (int): The number of successfully scanned and assigned items.
-> err (error): An error object if any parsing or conversion errors occur during the scan.
A non-nil value indicates an issue.
*/
