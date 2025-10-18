package main

import (
	"fmt"
)

func welcomDisplay (){
	fmt.Println("Welcome our Application")
}

func getInput () (int,int) {

	fmt.Println("Enter the two Numbers ->> ")
	var num1 int
	var num2 int
    fmt.Scan(&num1,&num2)
	return num1,num2
}

func addTwoNumber(num1 int,num2 int) {
	sum := num1 + num2
	fmt.Println("The Result is ",sum)
}

func goodByeDisplay () {
	fmt.Println("Thank you!!")
	fmt.Println("Good Bye!!!!!!")
}

func main () {

	welcomDisplay()
	num1,num2 := getInput()
	addTwoNumber(num1,num2)
	goodByeDisplay()
}