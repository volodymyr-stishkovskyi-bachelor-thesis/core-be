package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/cron"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/handlers"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/redisclient"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/repositories"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/vector"
)

var logger = logrus.New()

func init() {
	file, err := os.OpenFile("timings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feUrl := os.Getenv("FE_URL")

		w.Header().Set("Access-Control-Allow-Origin", feUrl)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

	rdb := redisclient.NewClient()

	cron.StartScrapeCron(rdb)
	vector.Init()

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	router.HandleFunc("/queries", handlers.QueryHandler).Methods("POST")
	router.HandleFunc("/queries", handlers.GetUserQueriesHandler).Methods("GET")
	router.HandleFunc("/resume", handlers.ResumeHandler).Methods("GET")
	router.HandleFunc("/leetcode", handlers.LeetCodeHandler).Methods("GET")

	routerWithCORS := withCORS(router)

	logger.Info("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, routerWithCORS))
}
