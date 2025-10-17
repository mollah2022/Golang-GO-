package main

import (
	"fmt"
)

func main () {
	var number1 int
	var number2 int

	fmt.Println("Enter the first Numer - ")
	fmt.Scan(&number1)

	fmt.Println("Enter the Second Number - ")
	fmt.Scan(&number2)

	var sum int

	sum = number1 + number2

	fmt.Println(sum)
}