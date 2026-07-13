package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "no website provided\n")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Fprint(os.Stderr, "too many arguments provided\n")
		os.Exit(1)
	}

	rawURL := os.Args[1]

	const maxConcurrency = 3
	cfg, err := configure(rawURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", rawURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	for url := range cfg.pages {
		fmt.Printf("URL: %s sighted\n", url)
	}
}
