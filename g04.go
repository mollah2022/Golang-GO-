package main

import (
	"fmt"
) 

type User struct {
	Name string
	Id int
	Job string
	Salary float64
}

func disPlayValue( user User) {
    fmt.Println(user.Name)
	fmt.Println(user.Id,user.Job,user.Salary)
}

func main () {

   var ob1 User

   ob1.Name = "Sajib Ahmed"
   ob1.Id = 53
   ob1.Job = "Engineer"
   ob1.Salary = 45550.95

   fmt.Println(ob1)

   ob2 := User{
	Name: "Rakib",
	Id: 76,
	Job: "Engineer",
	Salary: 35600.63,
   }

   disPlayValue(ob2)
}