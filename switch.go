package main

import (
	"fmt"
)

func main () {

	var Day int

	fmt.Scan(&Day)

	switch Day {
	case  1:
		fmt.Println("Satturady")

		case  2:
		fmt.Println("Sunday")

		case  3:
		fmt.Println("Monday")

		case  4:
		fmt.Println("The")

		case  5:
		fmt.Println("WED")

		case 6:
		fmt.Println("Thus")

		case  7:
		fmt.Println("Fri")

		default:
			fmt.Println("No Day")
	} 
	

}