package routes

import (
	"net/http"

	"ecommerce/controllers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /users", http.HandlerFunc(controllers.GetUsers))
	mux.Handle("GET /user/", http.HandlerFunc(controllers.GetUserByID))
	mux.Handle("POST /create-user", http.HandlerFunc(controllers.CreateUser))
	return mux
}
