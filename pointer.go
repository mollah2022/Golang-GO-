package main

import (
	"fmt"
)

func main () {
	var num int
	num = 100
	fmt.Println(num)

	fmt.Println("Pointer Address -->> ")

	value :=&num
	fmt.Println(value)

	fmt.Println("Then Print Point value --->> ")

	newNum := *value
	fmt.Println(newNum)
}