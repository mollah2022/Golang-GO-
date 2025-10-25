package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int 
	Job string
	salary float32
}

func displayValue(user User){
    fmt.Println("Name : ",user.Name)
	fmt.Println("Age : ",user.Age)
	fmt.Println("Job : ",user.Job)
	fmt.Println("Salary : ",user.salary)
}

func main () {
	var user1 User
	var user2 User

	user1.Name = "Sajib"
	user1.Age = 23
	user1.Job = "Engineer"
	user1.salary = 20500.60

	user2.Name = "Habib"
	user2.Age = 30
	user2.Job = "Businessman"
	user2.salary = 50600.68

	displayValue(user1)
	displayValue(user2)
}