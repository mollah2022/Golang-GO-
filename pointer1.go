package main

import (
	"fmt"
) 

func displayArry( number *[5] int){
	fmt.Println("Print Array use in Function --->>> ",number)
}

func main () {
	 
	var array = [5] int {1,2,3,4,5}
	fmt.Println("Array Value Print",array)

	var newarrray = &array

	displayArry(newarrray)
}