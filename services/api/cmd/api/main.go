package main

import (
	"log"
	"net/http"

	"github.com/agency/api/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("api service listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
