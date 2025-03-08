package repositories

import (
	"context"
	"time"
)

type ScrapedData struct {
	ID        int       `json:"id"`
	Source    string    `json:"source"`
	URL       string    `json:"url"`
	RawText   string    `json:"raw_text"`
	CreatedAt time.Time `json:"created_at"`
}

func SaveScrapedData(source, url, rawText string) error {
	_, err := DB.Exec(context.Background(),
		`INSERT INTO scraped_data (source, url, raw_text) VALUES ($1, $2, $3)`,
		source, url, rawText)
	return err
}

func GetScrapedData() ([]ScrapedData, error) {
	rows, err := DB.Query(context.Background(),
		`SELECT id, source, url, raw_text, created_at FROM scraped_data ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ScrapedData
	for rows.Next() {
		var d ScrapedData
		err := rows.Scan(&d.ID, &d.Source, &d.URL, &d.RawText, &d.CreatedAt)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	return data, nil
}
