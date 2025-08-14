package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"ecommerce/db"
	"ecommerce/models"
	"ecommerce/util"
)

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.UserName, &u.Email, &u.Password); err != nil {
			util.SendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if len(users) == 0 {
		util.SendError(w, "No users found", http.StatusNotFound)
		return
	}

	util.SendData(w, users, http.StatusOK)
}

// GET /user/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	var u models.User
	err := db.DB.QueryRow("SELECT id, username, email, password FROM users WHERE id=?", id).
		Scan(&u.ID, &u.UserName, &u.Email, &u.Password)

	if err == sql.ErrNoRows {
		util.SendError(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, u, http.StatusOK)
}

// POST /create-user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if strings.TrimSpace(u.UserName) == "" || strings.TrimSpace(u.Email) == "" || strings.TrimSpace(u.Password) == "" {
		util.SendError(w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !validateEmail(u.Email) {
		util.SendError(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Insert into DB
	res, err := db.DB.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", u.UserName, u.Email, u.Password)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil || id == 0 {
		util.SendError(w, "User creation failed", http.StatusInternalServerError)
		return
	}

	u.ID = int(id)

	util.SendData(w, u, http.StatusCreated)
}

// validateEmail checks if email format is valid
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
