package main

import "fmt"


func main () {

	var nums [5]int =[5]int{1,2,3,4,5}
	fmt.Println(nums)

	array := [6]int{1,2,3,4,5,6} 
	fmt.Println(array)

	array1 := array
	fmt.Println(array1)
	array1[1] = 100
	fmt.Println("------======------")
	fmt.Println(array1)

	for i:=0;i<len(array1);i++ {
		fmt.Println(array1[i])
	}

	fmt.Println("--------=========------------") 
	for  index,value := range nums {
		fmt.Println(index," == ",value)
	}
}