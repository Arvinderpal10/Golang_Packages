package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age : ")
	fmt.Scanln(&name, &age)
	fmt.Printf("Input : Name: %s, Age: %d\n", name, age)
}

/*
Scanln in Go is a function from the fmt package used to read space-separated values from standard
input, stopping at a newline. It is commonly used to read input from the user in a simple and
straightforward manner.


-> fmt.Scanln reads successive space-separated values from standard input and stores them in
successive arguments.
-> Input reading stops when a newline is encountered.
-> Each value must be on the same line and separated by spaces.
-> It returns the number of items successfully scanned and an error, if any.
*/
