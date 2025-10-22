package main

import (
	"fmt"
)

var (
	a = 10
	b = 20
)

func add(a int , b int ) {
	sum := a + b
	fmt.Println(sum)
}

func init () {
	fmt.Println("Hello Init Function-->>")
}

func main () {
	add(a,b)
}