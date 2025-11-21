package main

import (
	"log"
	"net/http"

	"github.com/RealImage/team-48gb/internal/router"
)

func main() {

	// Create router
	r := router.NewRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
