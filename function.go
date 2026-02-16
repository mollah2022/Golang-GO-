package main

import "fmt"


func add(n int,m int) {
	sum := n + m
	fmt.Println(sum)
}

func returnValue(n,m int) int {
	return n+m
}

func namedReturn(n,m int) (res int) {
	res = n + m 
	return res
}

func varaidicFunc(numbers ...int) {
	sum := 0

	for _,val:= range numbers {
		sum += val
	}

	fmt.Println(sum)
}

func higherOrderFunc(n,m int, value func(a,b int)){
	value(n,m)
}

func mul (n,m int) {
	sum := n * m
	fmt.Println("FirstOrder Func --->>>>> ",sum)
}


func main() {
	var num1,num2 int 

	func () {
		num1 = 10
	num2 = 30

	add(num1,num2)
	value := returnValue(num1,num2)
	fmt.Println(value)

	rValue := namedReturn(num1,num2)
	fmt.Println(rValue)
	}()

	varaidicFunc(1,2,3,4,5,6,7,8,9)
	higherOrderFunc(10,1000,mul)

}