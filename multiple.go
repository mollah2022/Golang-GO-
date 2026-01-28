package main

import "fmt"


func displayValue(a int,b int) (int,int) {
	sum := a + b
	mul := a - b

	return sum,mul
}


func main() {

	var num_one int
	var num_two int 

	fmt.Scanln(&num_one)
	fmt.Scanln(&num_two)
   
	sum,mul:= displayValue(num_one,num_two)

	fmt.Println("Sum ",sum,"Mal ",mul)

}