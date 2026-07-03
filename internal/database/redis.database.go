package database

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(uri string)( *redis.Client, error){
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second);
	defer cancel()

	options, err:= redis.ParseURL(uri)
	if err != nil {
		return nil , err
	}
	client := redis.NewClient(options)

	err = client.Ping(ctx).Err()

	if err != nil{
		return nil ,err
	}
	RedisClient = client
	fmt.Println("Redis Server is Connected")

	return client , nil
}