package main

import (
	"fmt"
)

func main () {

	fmt.Println("print String Value - ")

	name := "Sajib"

	fmt.Println("My name is ",name)

	var fName string
	var lName string

	fmt.Println("Enter the Two String Value -> ")
	fmt.Scan(&fName,&lName)

	fmt.Println("My Full Name is ",fName," ",lName)
}