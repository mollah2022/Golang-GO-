package main

import (
	"fmt"
)

func main () {
	var number int 
	number = 20
	fmt.Println(number)
	fmt.Println("Please Enter The Number --->>>> ")
	var num1 int
	fmt.Scan(&num1)
	fmt.Println("The number is ", num1)
}