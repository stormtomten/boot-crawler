package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read document: %w", err)
	}
	images := []string{}

	doc.Find("img[src]").Each((func(i int, s *goquery.Selection) {
		val, exists := s.Attr("src")
		if exists {
			trimmedVal := strings.TrimSpace(val)
			parsedURL, err := url.Parse(trimmedVal)
			if err == nil {
				resolvedURL := baseURL.ResolveReference(parsedURL)
				images = append(images, resolvedURL.String())
			}
		}
	}))

	return images, nil
}
