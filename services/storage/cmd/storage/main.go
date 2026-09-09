package main

import (
	"log"
	"net/http"

	"github.com/agency/storage/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("storage service listening on :8082")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
