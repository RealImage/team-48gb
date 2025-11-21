package main

import (
	"log"
	"net/http"

	"github.com/RealImage/team-48gb/internal/db"
	"github.com/RealImage/team-48gb/internal/router"
)

func main() {
	// Connect to MongoDB
	mongoURI := "mongodb://127.0.0.1:27017/?directConnection=true"
	dbName := "myTestDb"

	client, err := db.NewMongoClient(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect()

	// Create router with MongoDB client
	r := router.NewRouter(client)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
