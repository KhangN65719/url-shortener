package store

import (
	"sync"
)

type Data struct {
	mu     sync.RWMutex
	urlMap map[string]string
}

func New() *Data {
	return &Data{
		urlMap: make(map[string]string),
	}
}

func (data *Data) Write(code, longUrl string) {
	data.mu.Lock()
	defer data.mu.Unlock()
	data.urlMap[code] = longUrl
}

func (data *Data) Read(code string) string {
	data.mu.RLock()
	defer data.mu.RUnlock()
	return data.urlMap[code]
}
