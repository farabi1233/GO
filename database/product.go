package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"` // JSON tag for image URL
}

func init() {
	// Initialize some sample products
	prd1 := Product{
		ID: 1, Title: "Product 1", Description: "Description for product 1", Price: 19.99, ImageUrl: "http://example.com/image1.jpg",
	}
	prd2 := Product{
		ID: 2, Title: "Product 2", Description: "Description for product 2", Price: 29.99, ImageUrl: "http://example.com/image2.jpg",
	}
	prd3 := Product{
		ID: 3, Title: "Product 3", Description: "Description for product 3", Price: 39.99, ImageUrl: "http://example.com/image3.jpg",
	}
	ProductList = []Product{prd1, prd2, prd3} // Initialize the product list with sample products
}
