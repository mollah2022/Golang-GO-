package main

import (
	"fmt"
)

type Information struct {
	Name string
	Age int 
	Job string
	salary float32
}

func (user Information) displayValue1(){
    fmt.Println("Name : ",user.Name)
	fmt.Println("Age : ",user.Age)
	fmt.Println("Job : ",user.Job)
	fmt.Println("Salary : ",user.salary)
}

func main () {
	var info1 Information
	var info2 Information

	info1.Name = "Sajib"
	info1.Age = 23
	info1.Job = "Engineer"
	info1.salary = 20500.60

	info2.Name = "Habib"
	info2.Age = 30
	info2.Job = "Businessman"
	info2.salary = 50600.68

    info1.displayValue1()
	info2.displayValue1()
}