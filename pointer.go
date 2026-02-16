package main

import "fmt"


func hudai(m *int) {
	*m = 400
	fmt.Println("HUdai ----- >>>>> ",*m)
}

func main () {
	var number int 
	var poin *int 

	number = 100
	poin = &number
	fmt.Println(poin)
	fmt.Println(*poin)

	value := 200
	 p := &value
	fmt.Println(&value) 
	fmt.Println(*p)
	*p = 500
	fmt.Println(p)
	fmt.Println(*p)

	n := 800 
	 hudai(&n)
	fmt.Println("Check main ---->>>> ",n)
}