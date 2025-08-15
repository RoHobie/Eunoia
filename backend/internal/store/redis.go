package store

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Client *redis.Client

func ConnectRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis")
}

func GetRoomTTL() int {
	ttl, _ := strconv.Atoi(os.Getenv("ROOM_TTL_SECONDS"))
	if ttl <= 0 {
		ttl = 600
	}
	return ttl
}
