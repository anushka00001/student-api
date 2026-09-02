package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"student-api/config"
	"student-api/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	log.Println("GET /users")

	rows, err := config.DB.Query(
		"SELECT id, name, email FROM users",
	)

	if err != nil {
		log.Println("Database error:", err)

		http.Error(
			w,
			"Could not get users",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			http.Error(
				w,
				"Could not read user",
				http.StatusInternalServerError,
			)

			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {

	log.Println("POST /users")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	result, err := config.DB.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		user.Name,
		user.Email,
	)

	if err != nil {

		log.Println("Insert error:", err)

		http.Error(
			w,
			"Could not create user",
			http.StatusInternalServerError,
		)

		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		http.Error(
			w,
			"Could not get user ID",
			http.StatusInternalServerError,
		)

		return
	}

	user.ID = int(id)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	log.Println("PUT /users")

	id := r.URL.Query().Get("id")

	if id == "" {

		http.Error(
			w,
			"ID is required",
			http.StatusBadRequest,
		)

		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	result, err := config.DB.Exec(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		user.Name,
		user.Email,
		id,
	)

	if err != nil {

		log.Println("Update error:", err)

		http.Error(
			w,
			"Could not update user",
			http.StatusInternalServerError,
		)

		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(
			w,
			"Could not check update",
			http.StatusInternalServerError,
		)

		return
	}

	if rowsAffected == 0 {

		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)

		return
	}

	user.ID = 0

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	log.Println("DELETE /users")

	id := r.URL.Query().Get("id")

	if id == "" {

		http.Error(
			w,
			"ID is required",
			http.StatusBadRequest,
		)

		return
	}

	result, err := config.DB.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)

	if err != nil {

		log.Println("Delete error:", err)

		http.Error(
			w,
			"Could not delete user",
			http.StatusInternalServerError,
		)

		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(
			w,
			"Could not check delete",
			http.StatusInternalServerError,
		)

		return
	}

	if rowsAffected == 0 {

		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)

		return
	}

	w.Write([]byte("User deleted successfully"))
}
