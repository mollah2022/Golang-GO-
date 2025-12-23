package main

import "fmt"

func enterYourName(){
	fmt.Println("Please Enter Your Name ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("My name is ",name)
}

func addTwoNumber(n int, m int){
	res := n + m
	fmt.Println("Result",res)
}

func inputUserTwoNumber(){
	var num1 int
	fmt.Println("Please Enter First Number ")
	fmt.Scanln(&num1)
	var num2 int
	fmt.Println("Pleae Enter Second Number ")
	fmt.Scanln(&num2)
	addTwoNumber(num1,num2)
}

func main(){
	fmt.Println("Welcome! Get started with our application.")
	enterYourName()
	inputUserTwoNumber()
}