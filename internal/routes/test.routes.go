package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/handlers"
)

func TestRoutes(api *gin.RouterGroup) {
	test := api.Group("/test")
	test.GET("/health", handlers.Health)
}
