package model

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// Thread Safe connection pool
var RedisClient *redis.Client
var ctx = context.TODO()

func InitRedis() *redis.Client {

	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	RedisClient = redis.NewClient(options)

	pong, err := RedisClient.Ping(ctx).Result()

	if err != nil || pong != "PONG" {
		log.Fatal("Error connecting to Redis client")
	}

	fmt.Print("Connected to redis successfully \n")

	return RedisClient
}
