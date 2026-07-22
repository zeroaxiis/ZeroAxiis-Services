package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/handlers"
)

func AuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	auth.POST("/login", handlers.Login)
}
