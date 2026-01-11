package routes

import (
	"github.com/gin-gonic/gin"
)

// Setup returns a new gin.Engine with the routes for the service.
func Setup(r *gin.Engine) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Raspberry Pi Zero 2 W (emulated)!",
		})
	})

	// Ping endpoint
	// @Summary Ping the service
	// @Description This endpoint returns a simple "pong" message to test the service.
	// @Tags ping
	// @Accept json
	// @Produce json
	// @Success 200 {object} pongResponse
	// @Router /ping [get]
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, pongResponse{
			Message: "pong",
		})
	})

	// Welcome endpoint
	// @Summary Get a welcome message
	// @Description This endpoint returns a welcome message with information about the service.
	// @Tags welcome
	// @Accept json
	// @Produce json
	// @Success 200 {object} welcomeResponse
	// @Router /welcome [get]
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, welcomeResponse{
			Message: "Hello from Raspberry Pi Zero 2 W (emulated)!",
		})
	})

	// Health endpoint
	// @Summary Get the health status of the service
	// @Description This endpoint returns the health status of the service.
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} healthResponse
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, healthResponse{
			Status: "ok",
		})
	})

	return r
}

// pongResponse is the response object for the ping endpoint.
type pongResponse struct {
	Message string `json:"message"`
}

// welcomeResponse is the response object for the welcome endpoint.
type welcomeResponse struct {
	Message string `json:"message"`
}

// healthResponse is the response object for the health endpoint.
type healthResponse struct {
	Status string `json:"status"`
}
