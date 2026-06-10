package main

import (
	"github.com/KhangN65719/url-shortener/internal/handlers"
	"github.com/KhangN65719/url-shortener/internal/store"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	db := store.New()
	h := &handlers.Handler{Store: db}

	mux.HandleFunc("POST /shorten", h.Shorten)
	mux.HandleFunc("GET /{code}", h.Retrieve)

	err := http.ListenAndServe(":6767", mux)

	if err != nil {
		log.Fatal("Connection failed")
	}
}
