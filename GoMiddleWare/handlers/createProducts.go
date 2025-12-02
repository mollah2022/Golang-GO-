package handlers

import (
	"ecomerce/database"
	"ecomerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Products

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct) 

	if err != nil {

		fmt.Println(err)

		http.Error(w,"Please Give me a valid json",400)

		return

	}

	newProduct.ID = len(database.ProductsList) + 1

	database.ProductsList = append(database.ProductsList,newProduct)

	utils.SendData(w,database.ProductsList,201)

}
