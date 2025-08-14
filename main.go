package main

import (
	"ecommerce/cmd"
	"ecommerce/db"
)

func main() {
	db.Connect() // Connect to MySQL

	cmd.Serve()
}
