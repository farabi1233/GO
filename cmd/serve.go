package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_router"
	"ecommerce/handlers"
)

func Serve() {
	mux := http.NewServeMux() //route multiplexer
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /product/{productID}", http.HandlerFunc(handlers.GetProductByID))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	// User routes
	mux.Handle("GET /users", http.HandlerFunc(handlers.GetUsers))
	mux.Handle("GET /user/", http.HandlerFunc(handlers.GetUserByID))
	mux.Handle("POST /create-user", http.HandlerFunc(handlers.CreateUser))

	fmt.Println("Starting server on :3000")
	globalRouter := global_router.GlobalRouter(mux) // Apply global CORS middleware

	err := http.ListenAndServe(":3000", globalRouter) // start the server on port 3000
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
