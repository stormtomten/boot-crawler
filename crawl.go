package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v\n", err)
		return
	}

	if _, visited := pages[normalizedCurrent]; visited {
		pages[normalizedCurrent]++
		return
	}

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	pages[normalizedCurrent] = 1

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
		return
	}

	nextURLs, err := getURLsFromHTML(currentHTML, baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
		return
	}

	for _, url := range nextURLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
