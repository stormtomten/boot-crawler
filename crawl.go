package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v\n", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrent)
	if !isFirst {
		return
	}

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
		return
	}

	pageData := extractPageData(currentHTML, rawCurrentURL)
	cfg.setPageData(normalizedCurrent, pageData)

	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
