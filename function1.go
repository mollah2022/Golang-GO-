package main

import "fmt"


func displayResult(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	num1 := 10
	num2 := 20

	result := displayResult(num1,num2)
	fmt.Println(result)
}