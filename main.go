package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Products struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var ProductsList []Products

func getProducts(w http.ResponseWriter, r *http.Request){

	if r.Method != "GET" {
           http.Error(w,"Please send Get Request",400)
		   return
	}

	    w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(ProductsList)

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/products",getProducts)

	fmt.Println("Server Running Port is 8000 ")

	err := http.ListenAndServe(":8000",mux)

	if err!=nil {
		fmt.Println("Error Starting the server",err)
	}
}

func init(){
	pred1 := Products{
		Id: 1,
		Title: "Mango",
		Description: "Mango is my favorite Food. I love Mango",
		Price: 320,
		ImageUrl: "https://www.mango.org/wp-content/uploads/2024/06/Mango_Varieties_hero_img.png",
	}
	pred2 := Products{
		Id: 2,
		Title: "Banana",
		Description: "Banana is very Healthy Food. That is way i love Banana",
		Price: 120,
		ImageUrl: "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?fm=jpg&q=60&w=3000&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YmFuYW5hfGVufDB8fDB8fHww",
	}

	ProductsList = append(ProductsList, pred1)
	ProductsList = append(ProductsList, pred2)
}