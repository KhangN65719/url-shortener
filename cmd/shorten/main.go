package main

import (
	"fmt"
	"github.com/KhangN65719/url-shortener/internal/client"
	"github.com/KhangN65719/url-shortener/internal/validate"
	"log"
	"os"
)

func main() {
	url := os.Args[1]

	validURL, err := validate.ValidateLongURL(url)

	if err != nil {
		log.Fatal("Invalid URL")
	}

	if validURL && len(os.Args) == 2 {
		client.SendURL(url)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("No Arguments Found")
	}
}
