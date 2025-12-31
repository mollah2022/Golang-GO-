package main

import "fmt"


type list struct{
	name string
	age int
	dept string
	salary float32
}


func main(){

	var ob1 list

	ob1.name="Sajib"
	ob1.age = 25
	ob1.dept = "CSE"
	ob1.salary=200000

	fmt.Println(ob1)

	ob2 := list{
		name: "Rakib",
		age: 35,
		dept: "CSE",
		salary: 20300,
	}

	fmt.Println(ob2)

}