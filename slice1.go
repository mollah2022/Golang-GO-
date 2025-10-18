package main

import (
	"fmt"
)

func main () {

	var num int
	fmt.Scan(&num)

	slice := make([] int,0,num)

	for i :=0 ; i < num ; i++ {
		var x int
		fmt.Scan(&x)
		slice = append(slice,x)
	}

	for _ , v := range slice {
		fmt.Println(v)
	}

}