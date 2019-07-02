package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	hostname := os.Getenv("DB_HOSTNAME")

	DB, err = sql.Open("postgres",
		fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable", username, password, hostname, database))
	if err != nil {
		log.Fatal(err)
	}
}
