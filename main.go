package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello About Page")
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productlist []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please Give me Get Request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productlist)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/about", helloHandler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on port 3000...")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: " This is my favorite food",
		Price:       205,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSuQfOwRMG2Fk_6D880IWnkQf6BfkGEKjN7oQ&s",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: " This is my second favorite food",
		Price:       20,
		ImgUrl:      "https://img.freepik.com/free-photo/bananas-white-background_1187-1671.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Watermelon",
		Description: " This food is very testy",
		Price:       580,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRKZq-G705ZlMOJh69vQ4p3oEtDWogurMz2sQ&s",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Jackfruit",
		Description: " i do not like this food because is eating boring",
		Price:       930,
		ImgUrl:      "https://img.freepik.com/free-photo/yellow-jackfruit_74190-4803.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Lichi",
		Description: " I like this food because it is my fevorite food like ever",
		Price:       360,
		ImgUrl:      "https://www.kisanshop.in/uploads/muzaffarpur-lichi-plant.webp",
	}

	productlist = append(productlist, prd1)
	productlist = append(productlist, prd2)
	productlist = append(productlist, prd3)
	productlist = append(productlist, prd4)
	productlist = append(productlist, prd5)
}
