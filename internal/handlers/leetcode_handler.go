package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/redisclient"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/service"
)

func LeetCodeHandler(w http.ResponseWriter, r *http.Request) {
	raw, err := redisclient.Rdb.Get(r.Context(), "scrape:result").Result()
	if err == redis.Nil {
		http.Error(w, "LeetCode data not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Redis error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var scraped service.ScrapeResponse
	if err := json.Unmarshal([]byte(raw), &scraped); err != nil {
		http.Error(w, "Invalid cached data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var easy, medium, hard, totalEasy, totalMedium, totalHard int
	for _, s := range scraped.LeetCode.AcSubmissionNum {
		switch s.Difficulty {
		case "Easy":
			easy = s.Count
		case "Medium":
			medium = s.Count
		case "Hard":
			hard = s.Count
		}
	}
	for _, s := range scraped.LeetCode.AcSubmissionNum {
		switch s.Difficulty {
		case "Easy":
			totalEasy = 876
		case "Medium":
			totalMedium = 1840
		case "Hard":
			totalHard = 833
		}
	}

	resp := map[string]int{
		"easy":        easy,
		"totalEasy":   totalEasy,
		"medium":      medium,
		"totalMedium": totalMedium,
		"hard":        hard,
		"totalHard":   totalHard,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
