package main

import "fmt"

func main() {
	var number int

	fmt.Println("Please Enter The Number : ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Even Number")
	} else {
	     fmt.Println("Odd Number")	
	}

	var isValue int 

	fmt.Println("Please Enter The Value : ")
	fmt.Scanln(&isValue)

	switch isValue {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Error")
	}
}
