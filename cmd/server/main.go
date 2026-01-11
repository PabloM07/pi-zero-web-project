package main

import (
	"log"

	"context"

	"github.com/gin-gonic/gin"
	naas "github.com/pablom07/pi-gin-app/internal/naas"
	"github.com/pablom07/pi-gin-app/internal/routes"
)

// main is the entry point of the service.
// It sets up the routes and starts the server listening on port 8080.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	naasSvc := naas.NewService()
	naasSvc.Start(ctx)
	router := gin.Default()
	routes.Setup(router)
	routes.RegisterNaasRoutes(router, naasSvc)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
