package main

import (
	"fmt"
) 

func main () {

	var mapValue = map [string] string {"sajib":"Ahmed","Tamim":"Iqbal","Rashid":"Khan","Emran":"Tahir"}
	map1 := map[string] int {"sajib":1,"rakib":2,"tamim":3,"Shakib":23}

	fmt.Println(mapValue["sajib"])
    fmt.Println(map1)
}