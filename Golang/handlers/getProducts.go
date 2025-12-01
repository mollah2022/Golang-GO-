package handlers

import (
	"ecomerce/database"
	"ecomerce/utils"
	"net/http"
)



func GetProduct (w http.ResponseWriter,r *http.Request) {

	utils.SendData(w,database.ProductsList,200)
}


