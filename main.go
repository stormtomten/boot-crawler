package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Fprint(os.Stderr, "no website provided\n")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Fprint(os.Stderr, "too many arguments provided\n")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", args[0])
}
