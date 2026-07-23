package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeroaxiis/ZeroAxiis-Services/internal/pkg"
)

var RedisClient *redis.Client

func ConnectRedis(uri string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)

	err = client.Ping(ctx).Err()

	if err != nil {
		return nil, err
	}
	RedisClient = client
	pkg.Log.Info("Redis is connected")

	return client, nil
}
