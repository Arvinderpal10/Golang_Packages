package main

import (
	"fmt"
)

func main() {
	// Takes value till newline
	// Input : Arvinder 20               Output: Name : Arvinder , Age : 20
	// Input : Arvinder Pal 20           Output: Name : Arvinder , Age : 20
	// Input : Arvinder Pal              Output: Name : Arvinder , Age : 0
	// Input : 20 Arvinder               Output: Name : 20       , Age : 0
	// Input : 20 30                     Output: Name : 20       , Age : 30
	var name string
	var age int
	fmt.Print("Enter your name and age : ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Input by User : Name: %s, Age: %d\n", name, age)
}
