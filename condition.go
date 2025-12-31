package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	if num > 10 {
		fmt.Println("Hello Bangladesh")
	} else {
		fmt.Println("Hello Sajib")
	}
}