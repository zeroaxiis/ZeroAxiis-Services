package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/pkg"
	"go.uber.org/zap"
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
			pkg.Log.Error(err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})

			c.Abort()
			return
		}

		// Set Expiry on First Request - this will ensure when user created
		// first rquest then it will expire time to the user
		// note we are not sending TTL in first function bcz INCR do not support
		// 2 things it either take key or expiry at a time thats why we need and extra function
		if count == 1 {
			err = database.RedisClient.Expire(ctx, key, window).Err()
			if err != nil {
				pkg.Log.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"Success": false,
					"message": "Internal Server error",
				})
				c.Abort()
				return
			}
		}

		if count > limit {
			pkg.Log.Warn("Rate Limit exceeded", zap.String("ip:", ip), zap.String("route:", route))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too many request. Please try again later",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
