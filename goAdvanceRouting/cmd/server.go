package cmd

import (
	globalRouter "ecomerce/globalRouter"
	"ecomerce/handlers"
	"fmt"
	"net/http"
)

func Server () {

	mux := http.NewServeMux()

	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProduct))

	mux.Handle("POST /products",http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{id}",http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server Running on : 8080")

	globalRouter := globalRouter.GlobalRouter(mux)

	err:= http.ListenAndServe(":8080",globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}

}