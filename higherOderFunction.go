package main

import "fmt"

func addResult(x int, y int, fun func(a int, b int)){
	fun(x,y)
}

func callback() func(n int, m int){
	return add
}

func add(n int, m int){
	res := n + m
	fmt.Println("result",res)
}


func main(){
	
	fmt.Println("Welcome Explore our application")
	fmt.Println("Enter Two Number ")
	var num1,num2 int
	fmt.Scanln(&num1,&num2)
	addResult(num1,num2,add)

	res := callback()

	res(3,4)
}