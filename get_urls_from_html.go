package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read document: %w", err)
	}
	links := []string{}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			trimmedVal := strings.TrimSpace(val)
			parsedURL, err := url.Parse(trimmedVal)
			if err == nil {
				resolvedURL := baseURL.ResolveReference(parsedURL)
				links = append(links, resolvedURL.String())
			}
		}
	})

	return links, nil
}
