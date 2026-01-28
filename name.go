package main

import "fmt"


func display(a int,  b int) (res int) {
	res  = a + b

	return
}


func main() {
	var num int 
	var num1 int

	num = 10
	num1 = 39

	fmt.Println(display(num,num1))
}