package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/config"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/database"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/models"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/pkg"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/utils"
	"go.uber.org/zap"
)

// Handler for Halomail it -_- receives , verifies and saves event from halomail
func HanlderHalomailWebhook(c *gin.Context) {
	cfg := config.MustLoad()
	// this is for to verfiy sigature against header hehehehehehe (and it Blocks Fake request )
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		pkg.Log.Warn("Failed to read Halomail Webhook body", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.HalomailWebhookResponse{
			Success: false,
			Message: "Cannot read request body",
		})
		return
	}

	signature := c.GetHeader("X-HaloMail-Signature")
	if !utils.VerifyHaloMailSignature(rawBody, signature, cfg.HalomailWebhookSecret) {
		pkg.Log.Warn("Unauthorized Halomail Webhook : Invalid Signature")
		c.JSON(http.StatusUnauthorized, models.HalomailWebhookResponse{
			Success: false,
			Message: "Invalid or Missing Webhook Signature",
		})
		return
	}

	// Bind raw bytes into halomail request model struct -_+
	var request models.HalomailWebhookRequest
	if err := json.Unmarshal(rawBody, &request); err != nil {
		pkg.Log.Warn("Invalid JSON Payload in halomail webhook", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.HalomailWebhookResponse{
			Success: false,
			Message: "Malformed JSON Payload",
		})
		return
	}
	//Tranform Request into database, Finalllllllllllyyyyyyyy
	// this is for local data entry once done we will put it into database using database connection
	eventEntity := models.HaloMailEvent{
		Event:      request.Event,
		Timestamp:  request.Timestamp,
		Data:       request.Data,
		ReceivedAt: time.Now().UTC(),
	}

	collection := database.MongoClient.Database("Zeroaxiiscom").Collection("halomail_events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, eventEntity)
	if err != nil {
		pkg.Log.Error("Failed to insert Halomail event to database", zap.Error(err))
		c.JSON(http.StatusInternalServerError, models.HalomailWebhookResponse{
			Success: false,
			Message: "failed to persist event",
		})
		return
	}

	pkg.Log.Info("Halomail webhook Processed",zap.String("event",eventEntity.Event))
	c.JSON(http.StatusOK, models.HalomailWebhookResponse{
		Success:true,
		Message: "Webhook processed Successfully",
	})
	

}




