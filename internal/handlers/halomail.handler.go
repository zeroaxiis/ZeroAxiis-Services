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
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

// HandleHaloMailWebhook receives, verifies, and saves events from HaloMail, hehehehehehe
func HandleHaloMailWebhook(c *gin.Context) {
	cfg := config.MustLoad()

	// 1. Read raw body bytes for HMAC signature check -_-
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		pkg.Log.Warn("Failed to read HaloMail webhook body", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.HalomailWebhookResponse{
			Success: false,
			Message: "Cannot read request body",
		})
		return
	}

	// 2. Verify signature against header (Blocks fake Postman requests) -_+
	signature := c.GetHeader("X-HaloMail-Signature")
	if !utils.VerifyHaloMailSignature(rawBody, signature, cfg.HalomailWebhookSecret) {
		pkg.Log.Warn("Unauthorized HaloMail webhook: invalid signature", zap.String("signature", signature))
		c.JSON(http.StatusUnauthorized, models.HalomailWebhookResponse{
			Success: false,
			Message: "Invalid or missing webhook signature",
		})
		return
	}

	// 3. Bind raw bytes into our Request Struct
	var request models.HalomailWebhookRequest
	if err := json.Unmarshal(rawBody, &request); err != nil {
		pkg.Log.Warn("Invalid JSON payload in HaloMail webhook", zap.Error(err))
		c.JSON(http.StatusBadRequest, models.HalomailWebhookResponse{
			Success: false,
			Message: "Malformed JSON payload",
		})
		return
	}

	// 4. Validate timestamp to prevent replay attacks
	eventTime, err := time.Parse(time.RFC3339, request.Timestamp)
	if err == nil {
		// If the event was created more than 5 minutes ago, reject it as a stale/replayed request
		if time.Since(eventTime) > 5*time.Minute || time.Until(eventTime) > 5*time.Minute {
			pkg.Log.Warn("Rejected expired webhook: potential replay attack", zap.String("timestamp", request.Timestamp))
			c.JSON(http.StatusBadRequest, models.HalomailWebhookResponse{
				Success: false,
				Message: "Webhook event has expired",
			})
			return
		}
	}

	// 5. Transform Request Struct -> Database Entity Struct
	eventEntity := models.HaloMailEvent{
		Event:      request.Event,
		Timestamp:  request.Timestamp,
		Data:       request.Data,
		ReceivedAt: time.Now().UTC(),
	}

	// 6. Store in MongoDB collection "halomail_events" finallllllllllyyyyyy
	collection := database.MongoClient.
		Database("zeroaxiiscom").
		Collection("halomail_events")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, eventEntity)
	if err != nil {
		pkg.Log.Error("Failed to insert HaloMail event to MongoDB", zap.Error(err))
		c.JSON(http.StatusInternalServerError, models.HalomailWebhookResponse{
			Success: false,
			Message: "Failed to persist event",
		})
		return
	}

	pkg.Log.Info("HaloMail webhook processed", zap.String("event", eventEntity.Event))

	// 7. Return typed 200 OK Response Struct
	c.JSON(http.StatusOK, models.HalomailWebhookResponse{
		Success: true,
		Message: "Webhook processed successfully",
	})
}

// GetHaloMailEvents returns the latest 50 events for your Next.js dashboard UI, yep
func GetHaloMailEvents(c *gin.Context) {
	collection := database.MongoClient.
		Database("zeroaxiiscom").
		Collection("halomail_events")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Sort by received_at descending (newest first), limit to 50
	findOptions := options.Find().
		SetSort(bson.D{{Key: "received_at", Value: -1}}).
		SetLimit(50)

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		pkg.Log.Error("Failed to fetch HaloMail events", zap.Error(err))
		c.JSON(http.StatusInternalServerError, models.HaloMailEventsResponse{
			Success: false,
			Count:   0,
			Data:    nil,
		})
		return
	}
	defer cursor.Close(ctx)

	var events []models.HaloMailEvent
	if err := cursor.All(ctx, &events); err != nil {
		pkg.Log.Error("Failed to decode HaloMail events", zap.Error(err))
		c.JSON(http.StatusInternalServerError, models.HaloMailEventsResponse{
			Success: false,
			Count:   0,
			Data:    nil,
		})
		return
	}

	// Return typed Response Struct with count and array
	c.JSON(http.StatusOK, models.HaloMailEventsResponse{
		Success: true,
		Count:   len(events),
		Data:    events,
	})
}

// finallllllllllyyyyyy webhook handler finshed 