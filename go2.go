package main

import (
	"fmt"
) 

func displayName() {
	var Name string
	fmt.Scan(&Name)
	fmt.Println("The Name is ---->>>> ",Name)
}

func displayPhone() {
	var Phone int64
	fmt.Scan(&Phone)
	fmt.Println("The Phone Number is ---->>>  ",Phone)
}

func displayEmail() {
	var Email string
	fmt.Scan(&Email)
	fmt.Println("The Email is ---->>>> ",Email)
}

func displayVarification(num1 int,num2 int,operation string){
	if operation == "+"{
		add := num1+num2
		fmt.Println("The result is ",add)
	} else if operation == "-"{
		minus := num1-num2
        fmt.Println("The result is ",minus)
	} else if operation == "*"{
		mal := num1*num2
        fmt.Println("The result is ",mal)
	} else if operation == "/"{
		div := num1/num2
        fmt.Println("The result is ",div)
	} else{
		fmt.Print("Error Varification.....------>>>>><<<<<<<--------")
	}
}

func main () {

	fmt.Println("Welcome our Application --->>>>>>")
	fmt.Println("Please Enter Your Name --->>>> ")
	displayName()
	fmt.Println("Please Enter Your Phone----->>> ")
	displayPhone()
	fmt.Println("Please enter your Email----->>>> ")
	displayEmail()
	fmt.Println("Check Verification ----->>>> ")

	var value1,value2 int
	var char string

	fmt.Println("Please Enter the Two Numer ---- >>>> ")
	fmt.Scan(&value1,&value2)
	fmt.Println("Please Enter the Operation Value ----->>>> ")
	fmt.Scan(&char)
	displayVarification(value1,value2,char)

}