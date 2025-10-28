package main

import (
	"fmt"
) 

type USer struct {
	Name string
	Age int
	Job string 
	Salary float32
}

func main () {
    
	obj1 := USer{
		Name: "Sajib",
		Age: 32,
		Job: "Engineer",
		Salary: 28060.85,
	}

	fmt.Println(obj1)

	randomValue := &obj1

	fmt.Println(randomValue)
	fmt.Println(randomValue.Name)
	fmt.Println(*randomValue)

}