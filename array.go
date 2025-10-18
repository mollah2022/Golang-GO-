package main

import (
	"fmt"
)

func main () {

	var numbers = [6] int {1,2,3,4,5,6}
	fmt.Println(numbers)

	var names = [5] string {"sajib","rakib","emran","ridwan","anik"}
	fmt.Println(names)

	for  i  := 0; i < 6; i++ {
		fmt.Println(i)
	}

}