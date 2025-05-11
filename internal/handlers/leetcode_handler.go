package handlers

import (
	"encoding/json"
	"net/http"

	grpcclient "github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/grpc"
)

// Define response structures for the LeetCode stats
type LeetCodeStatsResponse struct {
	Reputation      int32           `json:"reputation"`
	Ranking         int32           `json:"ranking"`
	AcSubmissionNum []SubmissionNum `json:"ac_submission_num"`
}

type SubmissionNum struct {
	Difficulty string `json:"difficulty"`
	Count      int32  `json:"count"`
}

func LeetCodeStatsHandler(w http.ResponseWriter, r *http.Request) {
	leetCodeClient, conn, err := grpcclient.ConnectToLeetCodeService()
	if err != nil {
		http.Error(w, "Failed to connect to LeetCode Service", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	resp, err := grpcclient.GetLeetCodeStats(leetCodeClient)
	if err != nil {
		http.Error(w, "Failed to get LeetCode stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Convert the gRPC response type to our response type
	acSubmissions := make([]SubmissionNum, len(resp.AcSubmissionNum))
	for i, sub := range resp.AcSubmissionNum {
		acSubmissions[i] = SubmissionNum{
			Difficulty: sub.Difficulty,
			Count:      sub.Count,
		}
	}

	json.NewEncoder(w).Encode(LeetCodeStatsResponse{
		Reputation:      resp.Reputation,
		Ranking:         resp.Ranking,
		AcSubmissionNum: acSubmissions,
	})
}
