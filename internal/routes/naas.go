package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablom07/pi-gin-app/internal/naas"
)

type toggleRequest struct {
	Enabled bool `json:"enabled"`
}

func RegisterNaasRoutes(r *gin.Engine, svc *naas.Service) {
	r.GET("/naas/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"enabled": svc.IsEnabled(),
		})
	})

	r.POST("/naas/enable", func(c *gin.Context) {
		var req toggleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Enabled {
			svc.Enable()
		} else {
			svc.Disable()
		}

		c.JSON(200, gin.H{
			"enabled": svc.IsEnabled(),
		})
	})
}
