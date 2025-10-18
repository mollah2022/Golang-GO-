package main

import (
	"fmt"
)

func main () {

	var a = map[string] string {"Name":"Sajib Ahmed","Age":"27","Dept":"CSE"}
	 b := map[string] int {"index1":1,"index2":2,"index3":3}

	fmt.Println(a)
	fmt.Println(b)

	var newa = make(map[string]string)
	newa["Brand"] = "BMW"
	newa["Color"] = "Black"
	newa["Year"] = "2000"

	fmt.Println(newa)

}