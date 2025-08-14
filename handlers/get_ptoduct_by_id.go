package handlers

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/util"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productID") // Extract productID from the URL query parameters
	id, err := strconv.Atoi(productId)    // Convert productID to an integer
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200) // Send the product data as a response
			return
		}
	}

	util.SendData(w, "Data not found", 404)

}
