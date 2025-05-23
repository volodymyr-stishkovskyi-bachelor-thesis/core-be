package repositories

import (
	"context"
	"log"
	"time"
)

type Query struct {
	ID        int       `json:"id"`
	Query     string    `json:"query"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
	Chat      string    `json:"chat"`
}

func SaveQueryWithChat(chatID, query, response string) error {
	_, err := DB.Exec(context.Background(),
		`INSERT INTO user_queries (chat, query, response) VALUES ($1, $2, $3)`,
		chatID, query, response,
	)
	return err
}

func GetUserQueries() ([]Query, error) {
	rows, err := DB.Query(context.Background(),
		`SELECT id, query, response, created_at, chat FROM user_queries ORDER BY created_at DESC`)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var queries []Query
	for rows.Next() {
		var q Query
		err := rows.Scan(&q.ID, &q.Query, &q.Response, &q.CreatedAt, &q.Chat)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		queries = append(queries, q)
	}

	if len(queries) == 0 {
		log.Println("No queries found in the database.")
	}

	return queries, nil
}
