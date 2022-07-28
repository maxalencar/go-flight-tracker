package main

import (
	"log"
	"os"

	"flight-tracker/internal/flighttracker"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := flighttracker.NewServer()
	srv.Init()

	// Start server
	if err := srv.Start(port); err != nil {
		log.Fatalf("error starting server. err: %v", err)
	}
}
