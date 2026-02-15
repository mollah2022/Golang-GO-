package main

import "fmt"

func main () {
	var nums []int 
	fmt.Println(nums)

	num := make([]int,3)
	fmt.Println(num)
	num = append(num, 1)
	fmt.Println(num)

	for i:=1;i<=3;i++{
		nums = append(nums, i)
	}

	for _,val:= range nums {
		fmt.Println(val)
	}

	value := make([]int,0,6)
	fmt.Println(len(value),cap(value),value)

	for i:=1;i<=6;i++{
		value = append(value, i)
	}
	fmt.Println(value)

	for _,value:= range value {
		fmt.Println(value)
	}
}