package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/handlers"
)

func HealthRoutes(api *gin.RouterGroup) {
	health := api.Group("/quality")
	health.GET("/health", handlers.HealthHandler)
}