package main

import (
	"fmt"
) 

func main () {

	var array = [5] int {1,2,3}
	var array1 = [5] int {}
	var array2 = [5] int {1,2,3,4,5}

	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(array2)

	var value = [...] int {1,2,3,4,5,6,7,8}
	length := len(value)

	for i:=0 ; i < length ; i++ {
		fmt.Println(value[i])
	}
 
   marks := [...] int {1,2,3,4,5,6}
   marks[2] = 100

   fmt.Println(marks)

}