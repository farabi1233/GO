package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"ecommerce/db"
	"ecommerce/util"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // in production, don’t return password
}

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.UserName, &u.Email, &u.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	util.SendData(w, users, 200) // Send the product data as a response

}

// GET /user/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/user/")

	var u User
	err := db.DB.QueryRow("SELECT id, username, email, password FROM users WHERE id=?", id).
		Scan(&u.ID, &u.UserName, &u.Email, &u.Password)

	if err == sql.ErrNoRows {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, u, 200) // Send the product data as a response

}

// POST /create-user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if strings.TrimSpace(u.UserName) == "" || strings.TrimSpace(u.Email) == "" || strings.TrimSpace(u.Password) == "" {
		util.SendData(w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !validateEmail(u.Email) {
		util.SendData(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", u.UserName, u.Email, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	u.ID = int(id)

	util.SendData(w, u, 201) // Send the product data as a response

}

// validateEmail checks if email format is valid
func validateEmail(email string) bool {
	// Simple regex for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
