package main

import (
	"github.com/KhangN65719/url-shortener/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handlers.Shorten)
	mux.HandleFunc("GET /{code}", handlers.Retrieve)

	err := http.ListenAndServe(":6767", mux)

	if err != nil {
		log.Fatal("Connection failed")
	}
}
