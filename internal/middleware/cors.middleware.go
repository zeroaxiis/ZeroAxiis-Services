package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(publicFrontend string, adminFrontend string) gin.HandlerFunc {

	config := cors.Config{
		AllowOrigins: []string{
			publicFrontend,
			adminFrontend,
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}

	return cors.New(config)
}
