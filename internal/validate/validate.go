package validate

import (
	"fmt"
	"net/url"
)

func ValidateLongURL(raw string) (bool, error) {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false, fmt.Errorf("invalid URL: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false, fmt.Errorf("URL must start with http:// or https://")
	}
	if u.Host == "" {
		return false, fmt.Errorf("URL must have a valid host")
	}
	return true, nil
}
