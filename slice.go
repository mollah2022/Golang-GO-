package main

import "fmt"


func main(){ 
	var name = []int{1,2,3,4,5}

	fmt.Println(name)

	num := []int{1,2,3,4,5}

	fmt.Println(num)

	var value = make([]int,5,9)
	value[0] = 1
	fmt.Println(value)

	var n int
	fmt.Scanln(&n)

	var numbers = make([]int,0,n)

	for i:=0;i<n;i++{
		var v int
		fmt.Scanln(&v)
		numbers = append(numbers, v)
	}

		for i:=0; i<n;i++{
		fmt.Println(numbers[i])
	}
}