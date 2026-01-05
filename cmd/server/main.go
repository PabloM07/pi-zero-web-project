package main

import (
	"log"

	"github.com/pablom07/pi-gin-app/internal/routes"
)

// main is the entry point of the service.
// It sets up the routes and starts the server listening on port 8080.
func main() {
	r := routes.Setup()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
