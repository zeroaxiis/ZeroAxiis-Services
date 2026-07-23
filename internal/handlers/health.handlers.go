package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
)

func HealthHandler(c *gin.Context) {

	ctx := c.Request.Context()
	mongoHealthy := true

	err := database.MongoClient.Ping(ctx, nil)
	if err != nil {
		mongoHealthy = false
	}

	redisHealthy := true

	err = database.RedisClient.Ping(ctx).Err()
	if err != nil {
		redisHealthy = false
	}

	healthy := mongoHealthy && redisHealthy

	if !healthy {

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"status":  "unhealthy",
			"services": gin.H{
				"mongodb": mongoHealthy,
				"redis":   redisHealthy,
			},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  "healthy",
		"services": gin.H{
			"mongodb": mongoHealthy,
			"redis":   redisHealthy,
		},
	})
}
