package middleware

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
)

func RateLimiter(limit int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Client IP
		ip := c.ClientIP()

		// Get Current Route
		route := c.FullPath()
		route = strings.TrimPrefix(route, "/")

		// Generate Redis Key
		key := "rate:" + route + ":" + ip

		// Get Request Context
		ctx := c.Request.Context()

		// Increment Request Counter
		count, err := database.RedisClient.Incr(ctx, key).Result()
		if err != nil {
			utils.Log.Error(err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})

			c.Abort()
			return
		}

		// Temporary (count will be used in next step)
		_ = count
		_ = limit
		_ = window

		c.Next()
	}
}