package routes

import (
	"net/http"

	"student-api/handlers"
)

func RegisterRoutes() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			handlers.GetUsers(w, r)

		case http.MethodPost:
			handlers.CreateUser(w, r)

		case http.MethodPut:
			handlers.UpdateUser(w, r)

		case http.MethodDelete:
			handlers.DeleteUser(w, r)

		default:
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})
}
