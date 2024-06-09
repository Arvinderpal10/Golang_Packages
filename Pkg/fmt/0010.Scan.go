package main

import "fmt"

func main() {
	// Taking Single Input --> Takes input till space
	// Input : Arvinder        Output : Arvinder
	// Input : Arvinder Pal    Output : Arvinder (Pal input afetr the s[ace is ignored)
	// var name string
	// fmt.Println("Enter you name :")
	// fmt.Scan(&name)
	// fmt.Println("Name : ", name)

	//Taking more than one input --> Waits until all the values are not supplied
	var a, b, c int
	fmt.Println("Enter three numbers you want to add :")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Total sum : ", a+b+c)
}
