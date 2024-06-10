package main

import "fmt"

func main() {
	var name string
	var age int
	var country string

	input := "Arvinder 25 India"

	n, err := fmt.Sscanf(input, "%s %d %s", &name, &age, &country)
	if err != nil {
		fmt.Print("Error :", err)
	} else {
		fmt.Printf("Total Items : %d -> My Name is %s and I am %d years old. I live in %s.", n, name, age, country)
	}

}
