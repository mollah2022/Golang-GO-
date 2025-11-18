package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Products struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productsList []Products

func handleGetProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Please Give me Request", 400)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productsList)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/products", handleGetProduct)

	fmt.Println("Runninig Server Port 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}

}

func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Mango",
		Description: "Mango is my favorite food.it is very testy food.",
		Price:       230,
		ImgUrl:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}

	prd2 := Products{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is a internation food.it is good for health",
		Price:       500,
		ImgUrl:      "https://media.istockphoto.com/id/183422512/photo/fresh-red-apples-on-white-background.jpg?s=612x612&w=0&k=20&c=OmfmcQLh3mwp4pFVQn9Sr2U8VCyIgGtV6fuaDmd3Yrk=",
	}

	prd3 := Products{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is a local food in our country. it is gain energy in my body",
		Price:       20,
		ImgUrl:      "https://www.shutterstock.com/image-photo/bananas-grapes-600nw-518328943.jpg",
	}

	prd4 := Products{
		ID:          4,
		Title:       "Orange",
		Description: "Orange is my most favorite fruit . it is colour Orange",
		Price:       503,
		ImgUrl:      "https://www.lifelinehealthcarebd.org/data/frontImages/gallery/product_image/Oranges-fruit.jpg",
	}

	prd5 := Products{
		ID:          5,
		Title:       "Papaya",
		Description: "This is ripe Papaya fruit.It is very testy Food",
		Price:       403,
		ImgUrl:      "https://thumbs.dreamstime.com/b/papaya-25285564.jpg",
	}

	prd6 := Products{
		ID:          6,
		Title:       "Watermellon",
		Description: "This Green fruit . it means outside Green and inside in Red",
		Price:       700,
		ImgUrl:      "https://centershealthcare.com/wp-content/uploads/2022/04/Watermelon.webp",
	}

	productsList = append(productsList, prd1)
	productsList = append(productsList, prd2)
	productsList = append(productsList, prd3)
	productsList = append(productsList, prd4)
	productsList = append(productsList, prd5)
	productsList = append(productsList, prd6)
}
