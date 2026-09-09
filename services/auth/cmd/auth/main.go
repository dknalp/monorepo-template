package main

import (
	"log"
	"net/http"

	"github.com/agency/auth/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("auth service listening on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
