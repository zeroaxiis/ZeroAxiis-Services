package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/config"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/routes"
)

func main() {

	// Load Configuration
	cfg := config.MustLoad()

	// Connect MongoDB
	mongoClient, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
  _= mongoClient
	// Create Gin Engine
	router := gin.Default()

	// Routes- grouping
	api := router.Group("/api/v1")

	//redis connection heheheh
	redisClient , err := database.ConnectRedis(cfg.RedisURI)

	if err!=nil {
		log.Fatal(err)
	}
	_= redisClient
	routes.TestRoutes(api)
	// Start Server
	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}