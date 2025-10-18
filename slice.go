package main

import ("fmt")

func main () {

	 var n int
    fmt.Print("কত সংখ্যা দিতে চাও: ")
    fmt.Scan(&n)

    nums := make([]int, 0, n) // খালি slice তৈরি
    fmt.Println("সংখ্যাগুলো দাও:")

    for i := 0; i < n; i++ {
        var x int
        fmt.Scan(&x)
        nums = append(nums, x) // slice এ যোগ করা
    }

    fmt.Println("তোমার Slice মানগুলো:")
    for _, v := range nums {
        fmt.Println(v)
    }
}