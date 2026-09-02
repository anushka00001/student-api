package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() error {

	log.Println("Connecting to MySQL...")

	var err error

	DB, err = sql.Open(
		"mysql",
		"root:Anu7667@tcp(127.0.0.1:3306)/student_api",
	)

	if err != nil {
		return err
	}

	err = DB.Ping()

	if err != nil {
		return err
	}

	log.Println("MySQL connected successfully")

	return nil
}
