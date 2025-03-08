package handlers

import (
	"encoding/json"
	"net/http"

	grpcclient "github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/grpc"
)

type ScrapeRequest struct {
	URL string `json:"url"`
}

type ScrapeResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	var req ScrapeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	scraperClient, conn, err := grpcclient.ConnectToScraperService()
	if err != nil {
		http.Error(w, "Failed to connect to Scraper Service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	resp, err := grpcclient.Scrape(req.URL, scraperClient)
	if err != nil {
		http.Error(w, "Failed to scrape", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ScrapeResponse{
		Title:   resp.Title,
		Content: resp.Content,
	})
}
