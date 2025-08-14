package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct) // Decode the incoming JSON to a Product struct
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(database.ProductList) + 1                   // Assign a new ID
	database.ProductList = append(database.ProductList, newProduct) // Add the new product to the list
	util.SendData(w, newProduct, http.StatusCreated)

}
