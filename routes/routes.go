package routes

import (
	"net/http"

	"ecommerce/controllers"
	"ecommerce/handlers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /product/{productID}", http.HandlerFunc(handlers.GetProductByID))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /users", http.HandlerFunc(controllers.GetUsers))
	mux.Handle("GET /user/", http.HandlerFunc(controllers.GetUserByID))
	mux.Handle("POST /create-user", http.HandlerFunc(controllers.CreateUser))
	return mux
}
