package handler

import (
	"net/http"

	"github.com/agency/go-api-lib/response"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", Health)
	mux.HandleFunc("POST /upload", Upload)
	mux.HandleFunc("GET /files", ListFiles)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func Upload(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusCreated, map[string]string{"message": "file uploaded"})
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, []string{})
}
