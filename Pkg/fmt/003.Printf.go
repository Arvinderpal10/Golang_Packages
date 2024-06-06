package main

import (
	"fmt"
)

func main() {
	var age int = 20
	var salary float64 = 60.005
	var name string = "Arvinder Pal"
	var yes bool = true

	fmt.Printf("Hi, My name is %s and I am %d years old. My salary is %f. Everything is %t.\n", name, age, salary, yes)
	fmt.Printf("My name is %v ", name)
	fmt.Printf("\n")
	fmt.Printf("--------------------\n")
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of salary: %T\n", salary)
	fmt.Printf("Type of yes: %T\n", yes)
}

/*
- Printf formats its arguments based on a format specifier string.
- The format specifier string contains verbs that define the format for each argument.
- Common format verbs include :
	%s - String
	%d - Decimal integer
	%f - Floating-point number
	%t - Boolean
	%v - Default format for any value
	%T - Type of the value
	%p - Pointer address
- Unlike Println, Printf does not automatically add spaces between arguments or append a newline at the end.
*/
