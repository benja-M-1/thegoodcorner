package main

import (
	"github.com/benja-M-1/thegoodcorner/handlers"
	"github.com/benja-M-1/thegoodcorner/models"
	"log"
	"net/http"
)

func main() {
	err := models.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/fizzbuzz", handlers.FizzbuzzHandler)

	log.Fatal(http.ListenAndServe(":80", nil))

	defer models.DB.Close()
}
