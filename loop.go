package main

import "fmt"


func main () {

	nums := 10

	for i:=0; i<nums;i++ {
		
		if i == 3 {
			break
		} 

		fmt.Println(i)
	}

	i := 2

	for i <= nums {
		fmt.Println(i) 
		i++
	}
}