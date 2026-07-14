package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v\n", err)
		os.Exit(1)
	}
	if maxConcurrency < 1 {
		fmt.Println("Error - maxConcurrency can't be less than 1")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error - maxPages: %v\n", err)
		os.Exit(1)
	}

	cfg, err := configure(rawURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", rawURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}

	if err = writeJSONReport(cfg.pages, "report.json"); err != nil {
		fmt.Printf("Error - json report: %v\n", err)
		os.Exit(1)
	}
}
