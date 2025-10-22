package main

import (
	"fmt"
)

func main () {

 var numbers = [5] int {1,2,3,4,5}
fmt.Println(numbers)

   var array = [5] int {}

   for i:=0;i < 5; i++ {
	  fmt.Scan(&array[i])
   }
  fmt.Println(array)
}