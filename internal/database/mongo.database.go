package database

import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
var Client *mongo.Client

func ConnectMongo(uri string) (*mongo.Client, error){
	ctx ,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client , err := mongo.Connect(options.Client().ApplyURI(uri))

	if err !=nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err !=nil {
		return nil ,err
	}

	Client = client
	fmt.Println("Mongodb server is connected")
	return client ,nil
}