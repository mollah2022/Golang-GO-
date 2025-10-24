package main

import (
	"fmt"
)

type Information struct {
	 name string
	 age int
	 job string
	 salary float32
}

func displayValue( info Information) {
	fmt.Println(info.name)
	fmt.Println(info.age)
    fmt.Println(info.job)
	fmt.Println(info.salary)
}

func main () {
	var info1 Information 
	var info2 Information 
	 
	info1.name = "Sajib"
	info1.age = 25
	info1.job = "engineer"
	info1.salary = 20050.56

	info2.name = "Habib"
	info2.age = 26
	info2.job = "engineer"
	info2.salary = 30500.56

	displayValue(info1)
	displayValue(info2)
}