package main

import (
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/benja-M-1/thegoodcorner/handlers"
	"github.com/benja-M-1/thegoodcorner/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var container = &app.Container{}

// Initialize the applications and the dependency injection container
//
// It will
// - load .env file
// - connect to the database and ensure the connection will be closed once the program ends
// - register the http handlers
func main() {
	loadEnv()

	db := initDatabase()
	defer db.Close()
	container.DB = db

	fb := handlers.NewFizzBuzzHandler(container)
	http.HandleFunc("/fizzbuzz", fb.Handle)
	//http.HandleFunc("/stats", handlers.StatisticsHandler)

	log.Fatal(http.ListenAndServe(":80", nil))

}

// Create the connection to the database and returns it
func initDatabase() *models.DB {
	db, err := models.NewDB(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_HOSTNAME"),
		"5432",
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Read the .env file existing in the project and load them into ENV for this process.
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
