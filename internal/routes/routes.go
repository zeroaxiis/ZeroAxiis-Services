package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(api *gin.RouterGroup) {
	TestRoutes(api)
	HealthRoutes(api)
	AuthRoutes(api)
	TeamRoutes(api)
	TestimonialRoutes(api)
	BlogRoutes(api)
	CreativeRoutes(api)
	ProjectRoutes(api)
	HaloMailRoutes(api)
}
