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

	html, err := getHTML(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %v", err)
		os.Exit(1)
	}

	fmt.Println(html)
}
