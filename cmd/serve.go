package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_router"
	"ecommerce/routes"
)

func Serve() {
	mux := routes.RegisterRoutes()

	fmt.Println("Starting server on :3000")
	globalRouter := global_router.GlobalRouter(mux) // Apply global CORS middleware

	err := http.ListenAndServe(":3000", globalRouter) // start the server on port 3000
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
