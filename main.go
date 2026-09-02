package main

import (
	"log"
	"net/http"

	"student-api/config"
	"student-api/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Starting Student API...")

	err := config.ConnectDatabase()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	routes.RegisterRoutes()

	log.Println("Server running on http://localhost:4000")

	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("Server stopped:", err)
	}
}
