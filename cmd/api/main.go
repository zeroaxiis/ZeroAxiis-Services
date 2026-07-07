package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/config"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/middleware"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/routes"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
)

func main() {

	// Load Configuration
	cfg := config.MustLoad()

	//Logger initalization 
	err := utils.Init(cfg.AppEnv)
	if err !=nil{
		log.Fatal(err)
	}
	defer utils.Log.Sync()


	// Connect MongoDB
	mongoClient, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
  _= mongoClient


	// Create Gin Engine
	router := gin.Default()

	
	//cors setup
	router.Use(middleware.CORS(cfg.AdminFrontend, cfg.PublicFrontend))
	//security header hehehe
	router.Use(middleware.SecurityHeaders())


	// Routes- grouping
	api := router.Group("/api/v1")


	//redis connection heheheh
	redisClient , err := database.ConnectRedis(cfg.RedisURI)
	if err!=nil {
		log.Fatal(err)
	}
	_= redisClient


	//routes
	routes.TestRoutes(api)
	routes.HealthRoutes(api)


	// Start Server
	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}