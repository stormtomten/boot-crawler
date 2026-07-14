package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(url string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[url]; visited {
		return false
	}

	if len(cfg.pages) == cfg.maxPages {
		return false
	}

	cfg.pages[url] = PageData{}
	return true
}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	cfg.pages[normalizedURL] = data
	cfg.mu.Unlock()
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %v", err)
	}

	return &config{
		pages:              make(map[string]PageData),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
