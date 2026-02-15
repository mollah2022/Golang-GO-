package main

import "fmt"


func main () {

for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break
        }
        fmt.Println(i, j)
    }
}


   fmt.Println("       ------ =========== ---------- ")

    Outer:
	   for i := 0; i < 3; i++ {
		
		     for j := 0; j < 3; j++ {
				if i==1 && j==1 {
					break Outer
				}

				fmt.Println(i,j)
			 }
	   }

}