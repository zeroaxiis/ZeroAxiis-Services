package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type HaloMailEvent struct{
	ID bson.ObjectID `bson:"_id, omitempty" json:"id"`
	Event string `bson:"event" json:"event"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
	Data map[string] interface{} `bson:"data" json:"data"`
	ReceivedAt time.Time `bson:"received_at" json:"received_at"`
}

type HalomailWebhookRequest struct{
	Event string `json:"event" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
	Data map[string] interface{} `json:"data" binding:"required"`
}


type HalomailWebhookResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}


type HaloMailEventsResponse struct {
	Success bool            `json:"success"`
	Count   int             `json:"count"`
	Data    []HaloMailEvent `json:"data"`
}

