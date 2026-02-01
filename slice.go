package main

import "fmt"


func main() {
	
	nums := []int{1,2,3,55,66}
	fmt.Println(nums)
	nums[4] = 10
	for _,i := range nums {
		fmt.Println(i)
	}
}