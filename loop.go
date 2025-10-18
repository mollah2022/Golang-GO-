package main

import (
	"fmt"
)

func main () {

	var number int
	fmt.Scan(&number)

	for i :=1 ; i <= number ; i++ {
		if i % 2 == 0 {
			fmt.Println("Hello Even Ami -->> ",i)
		} else {
			fmt.Println("Break the Loop!!!!")
			break
		}
	}
}