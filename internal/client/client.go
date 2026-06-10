package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Link struct {
	URL string `json:"url"`
}

func SendURL(url string) {
	jsonInstance := Link{URL: url}
	jsonBody, err := json.Marshal(jsonInstance)

	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:6767/shorten", bytes.NewReader(jsonBody))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	fmt.Print(string(body))
}
