package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message":"The API is Working",
	})
}