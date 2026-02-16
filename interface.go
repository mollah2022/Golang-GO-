package main

import "fmt"


type User interface {
	getProduct()
}

type user struct {

}

func ( u user) getProduct() {
    fmt.Println("Hello BNP cholo chandra tuli")
}

func main() {

    a := user{
	}
	a.getProduct()

}