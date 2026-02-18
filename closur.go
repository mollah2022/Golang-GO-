package main

import "fmt"

const a = 10

var p = 100

func Outer() func() {
	money := 100
	age := 30

	fmt.Println("Age --->>>> ",age)

	show := func() {

		money = money + a + p
		fmt.Println(money)
	}

	return show
}
func call() {
	incr1 := Outer()
	incr1()
	incr1()

	incr2 := Outer()
	incr2()
	incr2()
}
func main() {
	call()
}
func init() {
	fmt.Println("Hello! I am Init Function....")
}