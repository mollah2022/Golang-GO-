package main

import "fmt"

func main() {
	
	var nums = [5]int{1,2,3,4,5}
	fmt.Println(nums)

	numbers := [5]int{1,2,3,4,5}
	fmt.Println(numbers)

	var arrayList  [8]string

     arrayList[0] = "sajib"
	 arrayList[1] = "rakib"
	 arrayList[2] = "tamim"
	 arrayList[3] = "habib"
	 arrayList[4] = "samim"
	 arrayList[5] = "fatema"
	 arrayList[6] = "yeamin"
	 arrayList[7] = "sumiya"

	 fmt.Println(arrayList)

	var numberList [5]int
	numberList[0] = 1
	numberList[1] = 2
	numberList[2] = 3
	numberList[3] = 4
	numberList[4] = 5

	fmt.Println(numberList)

	for i := 0; i < 5; i++ {
	fmt.Println(numberList[i])
}

    var n int
	fmt.Scanln(&n)

	var list [5]int

	for i:=0; i< n ;i++{
		fmt.Scanln(&list[i])
	}

	fmt.Println(list)

}