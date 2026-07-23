package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Admin struct{
	ID bson.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	LastLoginAt time.Time `bson:"lastLoginAt,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct{
	Message string `json:"message"`
	Admin AdminInfo `json:"admin"`
}

type AdminInfo struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`	
}

