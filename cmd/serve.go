package cmd

import (
	"fmt"
	"net/http"
	"os"

	"ecommerce/global_router"
	"ecommerce/routes"
)

func Serve() {
	mux := routes.RegisterRoutes()

	// Get the port from environment variable, default to 3000 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Starting server on :" + port)

	globalRouter := global_router.GlobalRouter(mux) // Apply global CORS middleware

	// Start the server
	err := http.ListenAndServe(":"+port, globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
