package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Format: username:password@tcp(host:port)/dbname
	dsn := "farabi:12345678@tcp(127.0.0.1:3306)/clients"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL successfully!")
}
