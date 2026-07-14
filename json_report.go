package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := []string{}
	for key := range pages {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sorted := []PageData{}

	for _, key := range keys {
		sorted = append(sorted, pages[key])
	}

	data, err := json.MarshalIndent(sorted, "", "  ")
	if err != nil {
		return fmt.Errorf("error - marshal report: %v", err)
	}

	if err = os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("error - writing file: %v", err)
	}

	return nil
}
