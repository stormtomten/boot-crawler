package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create reuest to: %s - %v", rawURL, err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to process request to: %s - %W", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("got HTTP status code %d for %s", resp.StatusCode, rawURL)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML content type %q for %s", contentType, rawURL)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %s - %w", rawURL, err)
	}

	return string(body), nil
}
