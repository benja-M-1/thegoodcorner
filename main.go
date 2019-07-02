package main

import (
	"github.com/benja-M-1/thegoodcorner/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fizzbuzz", handlers.FizzbuzzHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
