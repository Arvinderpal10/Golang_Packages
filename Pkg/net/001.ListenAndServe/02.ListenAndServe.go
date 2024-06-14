package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "----> Welcome to Home Page !!!!! ")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "----> Welcome to About Page !!!")
}
func main() {
	http.HandleFunc("/Home", homeHandler)
	http.HandleFunc("/About", aboutHandler)
	fmt.Println("Starting server on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Error in starting the server :", err)
	}
}
