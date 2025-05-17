package redisclient

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	// Rdb — глобальный клиент, доступен после первого вызова NewClient()
	Rdb  *redis.Client
	once sync.Once
)

// NewClient возвращает singleton Redis-клиент.
// Подключение и лог "Connected to Redis" будет выполнено только один раз.
func NewClient() *redis.Client {
	once.Do(func() {
		addr := os.Getenv("REDIS_URL")
		if addr == "" {
			log.Fatal("REDIS_URL environment variable is not set")
		}

		Rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDIS_PASSWORD"), // может быть пустым
			DB:       0,
		})

		if err := Rdb.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Redis connect failed: %v", err)
		}
		log.Println("Connected to Redis")
	})
	return Rdb
}
