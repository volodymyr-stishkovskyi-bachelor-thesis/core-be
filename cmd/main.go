package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/handlers"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
)

var logger = logrus.New()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = repositories.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	router.HandleFunc("/queries", handlers.SaveQueryHandler).Methods("POST")
	router.HandleFunc("/queries", handlers.GetUserQueriesHandler).Methods("GET")
	router.HandleFunc("/scraped", handlers.SaveScrapedDataHandler).Methods("POST")
	router.HandleFunc("/scraped", handlers.GetScrapedDataHandler).Methods("GET")

	logger.Info("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
