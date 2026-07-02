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

	// Create Gin Engine
	router := gin.Default()

	// // Register Routes
	routes.SetupRoutes(router, mongoClient)

	// Start Server
	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}