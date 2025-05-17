package cron

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/service"
)

func StartScrapeCron(rdb *redis.Client) {
	c := cron.New()
	// Cron every day at midnight
	c.AddFunc("0 0 * * *", func() {
		log.Println("Cron scraping")
		ctx := context.Background()
		resp, err := service.ConnectAndScrapeHandler()
		if err != nil {
			log.Println("  scrape error:", err)
			return
		}

		key := "scrape:result"
		oldData, err := rdb.Get(ctx, key).Bytes()
		if err == nil {
			var oldResp service.ScrapeResponse
			if err := json.Unmarshal(oldData, &oldResp); err == nil {
				if reflect.DeepEqual(oldResp, resp) {
					log.Println("  no changes since last scrape, skipping indexing")
					return
				}
			}
		}

		if err := service.IndexScrapeResponse(ctx, "scrape", resp); err != nil {
			log.Println("  indexing error:", err)
		} else {
			log.Println("  data indexed in Pinecone")
		}

		newData, _ := json.Marshal(resp)
		if err := rdb.Set(ctx, key, newData, 0).Err(); err != nil {
			log.Println("  redis set error:", err)
		} else {
			log.Println("  updated cache in Redis key=", key)
		}
	})
	c.Start()
	log.Println("Scrape cron started")
}
