package main

import (
	"fmt"
) 

func main () {
	var map1 = map [string] int {"index1":1,"index2":2,"index3":3,"index4":4}
	fmt.Println(map1)
	map1["index1"] = 100
	fmt.Println(map1)
	map1["index5"] = 200
	fmt.Println(map1)
	len1 := len(map1)
	fmt.Println(len1)
}          