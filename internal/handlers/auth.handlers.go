package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/models"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {

	var request models.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		utils.Log.Warn(
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
		utils.Log.Warn(
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
		utils.Log.Error("failed to find admin", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	err = utils.CheckPassword(admin.Password, request.Password)
	if err != nil {
		utils.Log.Warn(
			"login Failed Invalid Password",
			zap.String("email", request.Email),
		)
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid Email or Password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"message":"Login Successful",
	})

}
