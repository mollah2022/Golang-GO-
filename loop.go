package main

import "fmt"

func main() {

	num := 5
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	numbers := []int{10, 20, 30}

for i, v := range numbers {
	fmt.Println(i, v)
}

    m := map[string]int{
	"apple":  10,
	"banana": 20,
}

for key, value := range m {
	fmt.Println(key, value)
}


}