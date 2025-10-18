package main

import (
	"fmt"
)

func main () {

	fmt.Println("Prints Array -->> ")

	var numbers = [10] int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(numbers)

	var valuse = [] int {1,2,3}
	valuse = append(valuse, 20,39,90)

	fmt.Println("Prints Slice -->> ")

	fmt.Println(valuse)

	// how to get user input in array


	var arrayValue [10] int

	fmt.Println("Enter the insert array Value")

	for i :=0 ; i < 10 ; i++ {
		fmt.Scan(&arrayValue[i])
	}

	fmt.Println("Check the result in Array --> ")

	for i :=0 ; i < 10 ; i++ {
		fmt.Println(arrayValue[i])
	}
 

}