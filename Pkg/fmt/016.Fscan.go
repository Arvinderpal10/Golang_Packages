package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("016.Fscan.txt")
	if err != nil {
		fmt.Print("Error in creating file", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Arvinder 25 India")
	if err != nil {
		fmt.Print("Error in writing to a file")
		return
	}

	file, err = os.Open("016.Fscan.txt")
	if err != nil {
		fmt.Print("Error in opening a file")
		return
	}
	var name string
	var age int
	var country string

	n, err := fmt.Fscan(file, &name, &age, &country)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Total items %d. Name: %s, Age: %d, Country: %s\n", n, name, age, country)
	}

}

/*
-> Fscan in Go is a function from the fmt package used to read space-separated values from an io.Reader
and store them into successive arguments. This is useful for parsing structured data from sources
like files or network connections.


-> fmt.Fscan reads successive space-separated values from the provided io.Reader and stores
them in successive arguments.
-> It returns the number of items successfully scanned and an error, if any.
*/
