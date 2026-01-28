package main

import "fmt"

func sumValue(nums ...int) int {
	res := 0

	for _,n := range nums {
		res = res + n
	}

	return res
}


func main() {
   
	res := sumValue(1,2,3,4,5,6,7)

	fmt.Println(res)
}