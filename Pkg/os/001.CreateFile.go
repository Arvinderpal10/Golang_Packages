package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("001.data.txt")
	if err != nil {
		log.Fatalf("Failed to create a file : %s", err)
	}
	_, err = file.WriteString("This is my Second line of code.")
	if err != nil {
		log.Fatalf("Failed to create a file : %s ", err)
	}
	fmt.Println("File created and data written successfully.")
}
