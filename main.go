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

	fmt.Printf("starting crawl of: %s...\n", rawURL)

	pages := make(map[string]int)
	crawlPage(rawURL, rawURL, pages)

	for url, sightings := range pages {
		if sightings == 1 {
			fmt.Printf("URL: %s sighted %d time\n", url, sightings)
			continue
		}
		fmt.Printf("URL: %s sighted %d times\n", url, sightings)
	}
}
