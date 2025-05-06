package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("assets", "resume.pdf")
	feUrl := os.Getenv("FE_URL")

	w.Header().Set("Access-Control-Allow-Origin", feUrl)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"resume.pdf\"")

	http.ServeFile(w, r, path)
}
