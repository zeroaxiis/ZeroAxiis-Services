package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(api *gin.RouterGroup){
	TestRoutes(api)
	HealthRoutes(api)
}

