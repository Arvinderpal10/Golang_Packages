package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("017.Fscanf.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var name string
	var age int
	var country string

	n, err := fmt.Fscanf(file, "%s %d %s", &name, &age, &country)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	fmt.Printf("Successfully scanned %d items\n", n)
	fmt.Printf("Name: %s, Age: %d, Salary: %s \n", name, age, country)
}

/*
-> Fscanf function in Golang is used for formatted input. It reads formatted text from a
specified io.Reader (typically a file) and stores the extracted values into corresponding
variables.

-> Reads formatted data from a source.
-> Parses the data according to format specifiers.
-> Stores extracted values in provided variables.
n, err := fmt.Fscanf(r, format, a...)

Parameters:

-> r (mandatory): An io.Reader interface representing the source of the formatted text. This is usually a file opened using os.
Open or similar functions.
-> format (mandatory): A string specifying the format of the expected data. Similar to fmt.Printf format specifiers, it defines how to interpret the input text and match it with the provided variables.
-> a... (variadic): A sequence of variables where the extracted values will be stored. The number and types of variables must match the format specifiers in the format string.

Return Values:
-> n (int): The number of successfully scanned and assigned items.
-> err (error): An error object if any parsing or conversion errors occur during the scan.
A non-nil value indicates an issue.
*/
