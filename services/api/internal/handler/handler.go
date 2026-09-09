package handler

import (
	"net/http"

	"github.com/agency/go-api-lib/response"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", Health)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
