package handler

import (
	"net/http"

	"github.com/agency/go-api-lib/response"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", Health)
	mux.HandleFunc("POST /login", Login)
	mux.HandleFunc("POST /register", Register_)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"token": "placeholder"})
}

func Register_(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}
