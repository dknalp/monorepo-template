package main

import (
	"log"
	"net/http"
	"os"

	"github.com/agency/auth/internal/handler"
	"github.com/agency/go-api-lib/middleware"
)

func main() {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()
	handler.Register(mux)

	h := middleware.Logger(middleware.CORS(mux))

	log.Printf("auth service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal(err)
	}
}
