package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/openai"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/redisclient"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/vector"
)

var logger = logrus.New()

func init() {
	file, err := os.OpenFile("timings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func HandleQuery(ctx context.Context, chatID, query string) (string, error) {
	start := time.Now()
	emb, err := openai.GenerateEmbedding(ctx, query)
	if err != nil {
		return "", fmt.Errorf("embedding failed: %w", err)
	}
	logger.Infof("embed_ms=%d", time.Since(start).Milliseconds())

	pineStart := time.Now()
	matches, err := vector.QueryVectors(emb, 10)
	if err != nil {
		return "", fmt.Errorf("pinecone query failed: %w", err)
	}
	logger.Infof("pine_ms=%d", time.Since(pineStart).Milliseconds())

	rdb := redisclient.NewClient()
	raw, err := rdb.Get(ctx, "scrape:result").Result()
	if err != nil {
		return "", fmt.Errorf("redis get failed: %w", err)
	}
	var scraped ScrapeResponse
	if err := json.Unmarshal([]byte(raw), &scraped); err != nil {
		return "", fmt.Errorf("invalid scraped JSON: %w", err)
	}

	var docs []string
	for _, m := range matches {
		vec := m.Vector              // *pinecone.Vector
		id := vec.Id                 // string
		meta := vec.Metadata.AsMap() // map[string]interface{}

		typ, _ := meta["type"].(string)
		parts := strings.Split(id, ":")
		switch typ {
		case "credly":
			if len(parts) >= 3 {
				idx, _ := strconv.Atoi(parts[2])
				if idx >= 0 && idx < len(scraped.Credly) {
					c := scraped.Credly[idx]
					docs = append(docs,
						fmt.Sprintf("Certificate: %s; Issuer: %s; IssuedDate: %s",
							c.Title, c.Issuer, c.IssuedDate))
				}
			}

		case "leetcode":
			ll := scraped.LeetCode
			var easy, medium, hard int
			for _, s := range ll.AcSubmissionNum {
				switch s.Difficulty {
				case "Easy":
					easy = s.Count
				case "Medium":
					medium = s.Count
				case "Hard":
					hard = s.Count
				}
			}
			docs = append(docs,
				fmt.Sprintf("LeetCode stats — Reputation: %d; Ranking: %d; Easy: %d; Medium: %d; Hard: %d",
					ll.Reputation, ll.Ranking, easy, medium, hard))

		case "resume":
			if txt, ok := meta["text"].(string); ok && txt != "" {
				docs = append(docs,
					fmt.Sprintf("Resume excerpt: %s", txt))
			}
		}
	}

	var b strings.Builder
	b.WriteString("Use the following documents to answer the question:\n\n")
	for _, d := range docs {
		b.WriteString(d + "\n\n")
	}
	b.WriteString("Question: " + query)
	prompt := b.String()

	chatStart := time.Now()
	answer, err := openai.Chat(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("chat completion failed: %w", err)
	}
	logger.Infof("chat_ms=%d total_ms=%d",
		time.Since(chatStart).Milliseconds(),
		time.Since(start).Milliseconds())

	if err := repositories.SaveQueryWithChat(chatID, query, answer); err != nil {
		log.Printf("warning: failed to persist query: %v", err)
	}

	return answer, nil
}
