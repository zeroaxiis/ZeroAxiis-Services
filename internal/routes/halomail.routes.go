package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/handlers"
)

func HaloMailRoutes(api *gin.RouterGroup) {
	halomail := api.Group("/halomail")
	{
		// Outbound webhook receiver from HaloMail
		halomail.POST("/webhook", handlers.HandleHaloMailWebhook)

		// Next.js frontend feed
		halomail.GET("/events", handlers.GetHaloMailEvents)
	}
}