package main

import (
	"fmt"
)

func main () {
	
  var slice = [] int {1,2,3,4,5}
  var slice1 = [] string {"Rakib","Sajib","Habib","Rajbir"}

  fmt.Println(slice,slice1)

  number := 10

  var marks = make([] int,number)

  for i:=0; i < number ; i++ {
	fmt.Scan(&marks[i])
  }
  
  for i:=0 ; i < number; i++ {
	fmt.Println(marks[i])
  }

  var sliceValue = [] int {}

  sliceValue = append(slice,marks...)

  fmt.Println(sliceValue)

}