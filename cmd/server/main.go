package main

import (
	"log"
	"net/http"
	"github.com/KhangN65719/url-shortener/internal/handlers"
)

func main(){
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /shorten", handlers.Shorten)
	mux.HandleFunc("GET /{code}", handlers.Retrieve)
		
	err := http.ListenAndServe(":6767", mux)

	if err != nil{
		log.Fatal("Connection failed")
	}
}
