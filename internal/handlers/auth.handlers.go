package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/config"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/models"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/pkg"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {

	var request models.LoginRequest
	cfg := config.MustLoad()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		pkg.Log.Warn(
			"Invalid Login Request",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid Request",
		})
		return
	}
	adminCollection := database.MongoClient.Database("zeroaxiiscom").Collection("admin")

	var admin models.Admin

	err = adminCollection.FindOne(
		context.Background(),
		bson.M{
			"email": request.Email,
		},
	).Decode(&admin)

	if errors.Is(err, mongo.ErrNoDocuments) {
		pkg.Log.Warn(
			"login Failed: admin not found",
			zap.String("email", request.Email),
		)
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}
	if err != nil {
		pkg.Log.Error("failed to find admin", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	err = utils.CheckPassword(admin.Password, request.Password)
	if err != nil {
		pkg.Log.Warn(
			"login Failed Invalid Password",
			zap.String("email", request.Email),
		)
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid Email or Password",
		})
		return
	}
	sessionID, err := utils.CreateSession(admin.ID.Hex())
	if err != nil {
		pkg.Log.Error(
			"Failed to create session",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}
	token, err := utils.GenerateJWT(
		admin.ID.Hex(),
		sessionID,
		cfg.JWTSecret,
	)
	if err != nil {
		pkg.Log.Error(
			"Failed to generate JWT",
			zap.String("email", request.Email),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	c.SetCookie(
		"token",
		token,
		15*60,
		"/",
		"",
		false,
		true,
	)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login Successful",
	})

}
