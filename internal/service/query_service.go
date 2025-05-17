package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/openai"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/redisclient"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/vector"
)

// HandleQuery делает RAG + ChatGPT + сохраняет историю
func HandleQuery(ctx context.Context, chatID, query string) (string, error) {
	// 1) Эмбеддинг запроса
	emb, err := openai.GenerateEmbedding(ctx, query)
	if err != nil {
		return "", fmt.Errorf("embedding failed: %w", err)
	}

	// 2) Поиск в Pinecone (top-5)
	matches, err := vector.QueryVectors(emb, 10)
	if err != nil {
		return "", fmt.Errorf("pinecone query failed: %w", err)
	}

	rdb := redisclient.NewClient()
	// 3) Достаём последний скрейп из Redis
	raw, err := rdb.Get(ctx, "scrape:result").Result()
	if err != nil {
		return "", fmt.Errorf("redis get failed: %w", err)
	}
	var scraped ScrapeResponse
	if err := json.Unmarshal([]byte(raw), &scraped); err != nil {
		return "", fmt.Errorf("invalid scraped JSON: %w", err)
	}

	// 4) Собираем документы из совпадений по ID
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
			// в метаданных сохранили кусочек текста резюме
			if txt, ok := meta["text"].(string); ok && txt != "" {
				docs = append(docs,
					fmt.Sprintf("Resume excerpt: %s", txt))
			}
		}
	}

	// 5) Формируем промт
	var b strings.Builder
	b.WriteString("Use the following documents to answer the question:\n\n")
	for _, d := range docs {
		b.WriteString(d + "\n\n")
	}
	b.WriteString("Question: " + query)
	prompt := b.String()

	// 6) Запрашиваем ChatGPT
	answer, err := openai.Chat(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("chat completion failed: %w", err)
	}

	// 7) Сохраняем историю
	if err := repositories.SaveQueryWithChat(chatID, query, answer); err != nil {
		log.Printf("warning: failed to persist query: %v", err)
	}

	return answer, nil
}
