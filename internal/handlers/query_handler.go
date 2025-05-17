package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/service"
)

type QueryRequest struct {
	Chat  string `json:"chat"`
	Query string `json:"query"`
}

type QueryResponse struct {
	Response string `json:"response"`
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	var req QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	answer, err := service.HandleQuery(r.Context(), req.Chat, req.Query)
	if err != nil {
		http.Error(w, "Failed to process query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(QueryResponse{Response: answer})
}

func GetUserQueriesHandler(w http.ResponseWriter, r *http.Request) {
	queries, err := repositories.GetUserQueries()
	if err != nil {
		http.Error(w, "Failed to retrieve queries", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(queries)
}
