package main

import "fmt"

func main() {
	var nums [5]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 6

	fmt.Println(nums)



	var number = [5]int{1,2,3,4,5}
	fmt.Println(number)

	for index,value := range number {
		fmt.Println(index," ",value)
	}

	array := [5]int{1,23,4,6,7}
	fmt.Println(array)

	for _,i := range array {
		if(i == 6){
			break
		} else{
			fmt.Println(i)
		}
	}
}