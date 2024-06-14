package main

import (
	"errors"
	"fmt"
)

func divideByZero(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide %d by %d: %w", a, b, errors.New("division by zero"))
	}
	return a / b, nil
}
func main() {
	var a int = 50
	var b int = 0
	result, err := divideByZero(a, b)
	if err != nil {
		fmt.Print("Error :", err)
		return
	}
	fmt.Print("Result :", result)

}

/*
Errorf is a function in the fmt package in Go that formats according to a format specifier and
returns the string as a value that satisfies the error interface. This is useful for creating
formatted error messages in a standardized way.


-> Errorf formats its arguments based on a format specifier string, similar to Printf.
-> It returns an error type that contains the formatted string.
-> The returned error can be used in standard Go error handling.
*/
