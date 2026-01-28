package main

import "fmt"

func callbackFunction(a int, b int, op func(n int,m int) int ) int {
	return op(a,b)
}

func add(a int, b int) int {
	return a+b
}

func main() {
	callback := callbackFunction(3,4,add)
	fmt.Println(callback)
}