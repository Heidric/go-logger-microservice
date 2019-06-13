package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	// sslmode=verify-full
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", os.Getenv("DB_LOGIN"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect to DB with error:", err)
	} else {
		log.Println("Connected to DB")
	}
}
