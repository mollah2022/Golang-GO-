package cmd

import (
	globalRouter "ecomerce/globalRouter"
	"ecomerce/middleware"
	"fmt"
	"net/http"
)

func Server () {

	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Logger,middleware.Hudai)

	InitRoutes(mux,manager)

	fmt.Println("Server Running on : 8080")

	globalRouter := globalRouter.GlobalRouter(mux)

	err:= http.ListenAndServe(":8080",globalRouter)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}

}