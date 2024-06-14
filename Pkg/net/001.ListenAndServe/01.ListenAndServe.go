package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers !!!!!!! ")

}
func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server on 8080 port")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error in starting the server at : 8080", err)
	}

}
