package main

import (
	"fmt"
)

var (
	num1 = 20
	num2 = 30
)

func CallbackFunction(x int, y int, result func(a int, b int)) {
	result(x, y) 
}

func resultAllFunction(value int, value1 int) {
	sum := value + value1
	fmt.Println(sum)
}

func main() {
	CallbackFunction(num1, num2, resultAllFunction)
}
