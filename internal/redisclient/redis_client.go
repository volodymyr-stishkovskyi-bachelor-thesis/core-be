package redisclient

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb  *redis.Client
	once sync.Once
)

func NewClient() *redis.Client {
	once.Do(func() {
		addr := os.Getenv("REDIS_URL")
		if addr == "" {
			log.Fatal("REDIS_URL environment variable is not set")
		}

		Rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		if err := Rdb.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Redis connect failed: %v", err)
		}
		log.Println("Connected to Redis")
	})
	return Rdb
}
