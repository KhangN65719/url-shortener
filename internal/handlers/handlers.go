package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/KhangN65719/url-shortener/internal/store"
)

type Links struct {
	Url string `json:"url"`
}

type Handler struct {
	Store *store.Store
}

var codeLen uint8 = 6

func generateCode(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func (handler *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var myLink Links
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&myLink)

	existingCode := handler.Store.FindByLongUrl(myLink.Url)
	if existingCode != "" {
		fmt.Fprintf(w, "localhost:6767/%s\n", existingCode)
		return
	}

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	code := generateCode(int(codeLen))
	handler.Store.Write(code, myLink.Url)
	fmt.Fprintf(w, "localhost:6767/%s\n", code)
}

func (handler *Handler) Retrieve(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	originalUrl := handler.Store.Read(code)
	if originalUrl == "" {
		http.Error(w, "short link not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
