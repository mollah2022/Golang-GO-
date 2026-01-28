package main

import "fmt"

func result(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	var num0 int
	var num1 int

	num0 = 10
	num1 = 20

	result(num0, num1)
}