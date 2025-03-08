package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
)

func SaveScrapedDataHandler(w http.ResponseWriter, r *http.Request) {
	var data repositories.ScrapedData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := repositories.SaveScrapedData(data.Source, data.URL, data.RawText)
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetScrapedDataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := repositories.GetScrapedData()
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
