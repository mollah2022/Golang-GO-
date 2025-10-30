package main

import (
	"fmt"
)

func main () {
	
	var slice = [] int {1,2,3,4,5}
	fmt.Println(slice)
	slice = append(slice, 100,200)
	fmt.Println(slice)

	slice1 := make([] int,6,10)
	for i:=0;i<6; i++{
		fmt.Scan(&slice1[i])
	}
	fmt.Println(slice1)
}