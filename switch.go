package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	case 4:
		fmt.Println("FOUR")
	case 5:
		fmt.Println("FIVE")
	case 6:
		fmt.Println("SIX")
	case 7:
		fmt.Println("SEVEN")
	default:
		fmt.Println("Invalid number")
	}
}
