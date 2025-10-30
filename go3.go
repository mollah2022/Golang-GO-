package main

import (
	"fmt"
)

func displayArray(array [5] int ){
	fmt.Println("This Array " ,array)
}

func main () {
	var array = [5] int {1,2,3,4,5}
	fmt.Println(array)
	array1 := [6] int {1,2,3,4,5,6}
	fmt.Println(array1)

	array1[3] = 100
	array[3] = 200

	for i:=0 ; i < 5; i++ {
		fmt.Println(array[i])
	}

	for i:=0; i < len(array1) ; i++ {
		fmt.Println(array1[i])
	}

	displayArray(array)

}