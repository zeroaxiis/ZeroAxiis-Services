package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(publicFrontend string, adminFrontend string) gin.HandlerFunc {
	// Allowed production domains
	originMap := map[string]bool{
		"https://zeroaxiis.tech":       true,
		"https://dev.zeroaxiis.tech":   true,
		"https://admin.zeroaxiis.tech": true,
	}

	// Strictly allow localhost only in local development, never in production
	if os.Getenv("APP_ENV") != "production" {
		originMap["http://localhost:3000"] = true
		originMap["http://127.0.0.1:3000"] = true
	}

	// Add dynamic origins from environment variables (supports comma-separated lists)
	for _, raw := range []string{publicFrontend, adminFrontend} {
		for _, o := range strings.Split(raw, ",") {
			cleaned := strings.TrimSpace(strings.TrimSuffix(o, "/"))
			if cleaned != "" {
				originMap[cleaned] = true
			}
		}
	}

	allowedOrigins := make([]string, 0, len(originMap))
	for o := range originMap {
		allowedOrigins = append(allowedOrigins, o)
	}

	config := cors.Config{
		AllowOrigins: allowedOrigins,

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
		},

		ExposeHeaders: []string{
			"Content-Length",
			"Authorization",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}

	return cors.New(config)
}

