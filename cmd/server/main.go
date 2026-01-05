package main

import (
	"log"

	"github.com/pablom07/pi-gin-app/internal/routes"
)

func main() {
	r := routes.Setup()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
