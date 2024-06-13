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
