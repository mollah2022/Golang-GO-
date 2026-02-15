package main

import "fmt"



type User struct {
	Name string  `json:"name"`
	Email string `json:"email"`
	Password string `json:"passwor"`
	Age int `json:"age"`
}

var userList []User

func (u *User) display() {
	fmt.Println(u.Name)
}

func NewUser(name string,email string,pass string,age int) *User {
	return &User{
		Name: name,
		Email: email,
		Password: pass,
		Age: age,
	}
}


func main() {

	var user User

	user.Name = "Sajib Ahmed"
	user.Email = "sajib@gmail.com"
	user.Password = "123456"
	user.Age = 26 

	user.display()

	fmt.Println(user)

	user1 := User {
		Name: "Rakib",
		Email: "rakib@gmail.com",
		Password: "123456789",
		Age: 29,
	}
	user1.display()

	fmt.Println(user1)

	userList = append(userList, user)
	userList = append(userList, user1)

	fmt.Println("--------========----------")
	fmt.Println(userList)


	usr := NewUser("sajib","ahmed@gmail.com","151278",29)

	fmt.Println("**********===========***********")
	fmt.Println(usr)
    
	usr.display()

	userList = append(userList, *usr)
	fmt.Println("...........==========.........")
	fmt.Println(userList)

}