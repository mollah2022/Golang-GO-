package main

import (
	"fmt"
)

func main () {


	var mapValue = make(map[string] int)
	mapValue["sss"] = 1
	mapValue["aaa"] = 2
	mapValue["ccc"] = 3
	mapValue["ddd"] = 4

	fmt.Println(mapValue)

	for v,value:=range mapValue {
		fmt.Println(v,value)
	}

}