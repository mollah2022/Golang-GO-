package main

import "fmt"

type User interface {
	displayDetails()
}

type People interface {
	anotherDisplayDetails()
}

type user struct {
	Name     string
	Email    string
	Password string
	Age      int
}

func (obj user) displayDetails() {
	fmt.Println(obj.Name)
	fmt.Println(obj.Email)
	fmt.Println(obj.Password)
	fmt.Println(obj.Age)
}

func ( obj user) anotherDisplayDetails() {
	fmt.Println(obj.Name)
	fmt.Println(obj.Email)
	fmt.Println(obj.Password)
	fmt.Println(obj.Age)
}

func main() {
     
	 var usr1 user
	 usr1.Name = "Sajib Ahmed"
	 usr1.Email = "sajib@gmail.com"
	 usr1.Password = "151278"
	 usr1.Age = 28

	 user2 := user {
		Name: "Rakib Hossain",
		Email: "rakib@gmail.com",
		Password: "151278",
		Age: 29,
	 }

	 var user3 User
	 user3 = user{
		Name: "Habib Ahmed",
		Email: "habib@gmail.com",
		Password: "151278",
		Age: 22,
	 }

	 var user4 People
	 user4 = user{
		Name: "Akash Modak",
		Email: "akash@gmail.com",
		Password: "151278",
		Age: 25,
	 }

	 usr1.displayDetails()
	 user2.displayDetails()
	 user3.displayDetails()
	 user4.anotherDisplayDetails()

}