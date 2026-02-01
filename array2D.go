package main

import "fmt"

func main() {
    matrix1 := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    fmt.Println(matrix1)


        matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(matrix[i][j], " ")
        }
        fmt.Println()
    }

    for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Print(matrix[i][j], " ")
    }
    fmt.Println()
}



}
