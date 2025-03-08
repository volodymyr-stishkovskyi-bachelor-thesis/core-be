package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
)

func SaveQueryHandler(w http.ResponseWriter, r *http.Request) {
	var query repositories.Query
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := repositories.SaveQuery(query.Query, query.Response)
	if err != nil {
		http.Error(w, "Failed to save query", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUserQueriesHandler(w http.ResponseWriter, r *http.Request) {
	queries, err := repositories.GetUserQueries()
	if err != nil {
		http.Error(w, "Failed to retrieve queries", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(queries)
}
