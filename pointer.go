package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	*b = 20

	fmt.Println(a)
	fmt.Println(*b)
}