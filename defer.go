package main

import "fmt"

func calculator() (result int){
    
	fmt.Println("Calculator First Result ---->>>>> ",result)
   defer func() {
		result = result + 10
		fmt.Println("Defer Calculator Result ---->>>>>> ",result)
	}()

	result = result + 5
	fmt.Println("Calculator Second Result ----->>>>>> ",result)
	return 
}
func Callc() int{
    result := 0 
	fmt.Println("Callc First Result ---->>>>> ",result)

  	defer func() {
		result = result + 10
		fmt.Println("Defer Callc Result ---->>>>>> ",result)
	}()
	result = result + 5
	fmt.Println("Callc Second Result ----->>>>>> ",result)

	return result
}
func main() {
	
	calculator := calculator()
	fmt.Println("Main calculator------->>>>>>>>>> ",calculator)

	callc := Callc()
	fmt.Println("Main Callc------->>>>>>>>>>",callc)
} 