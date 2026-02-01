package main

import "fmt"


func main() {

	var number1 []int
	number1 = make([]int,3,5)
	fmt.Println(number1)

	for i:=0;i<5;i++{
		
		var value int
		fmt.Scanln(&value)
        number1 = append(number1, value)
	}

	fmt.Println(number1)


	slice := []int{}

	for i:=0;i<5;i++ {
		slice = append(slice,i)
	}

	for _,i:= range slice {
		fmt.Println(i)
	}

}