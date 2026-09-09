package main

import (
	"log"
	"net/http"
	"os"

	"github.com/agency/api/internal/handler"
	"github.com/agency/go-api-lib/middleware"
)

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	handler.Register(mux)

	h := middleware.Logger(middleware.CORS(mux))

	log.Printf("api service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal(err)
	}
}
