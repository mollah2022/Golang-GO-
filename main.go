package main

import (
	"example/calcul"
	"fmt"
)

func printName(){
	fmt.Println("Please Enter Your Name ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("My Name is ",name)
}

func twoNumber(){

	var n int 
	var m int

	fmt.Println("Please Enter first Number")
	fmt.Scanln(&n)

	fmt.Println("Please Enter Second Number")
	fmt.Scanln(&m)

	calcul.Add(n,m)

}


func main(){

	fmt.Println("Welcome our Website")

	printName()

	twoNumber()
}