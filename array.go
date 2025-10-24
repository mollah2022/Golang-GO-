package main

import (
	"fmt"
)

func main () {

	var array1 = [5] int {1,2,3,4,5}

	var array = [5] int {}

	for i:=0 ; i < 5; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println(array1)

	fmt.Println(array1)
}