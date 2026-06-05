package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "Basic one image relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:     "Basic two images relative",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			<img src="/logo.png" alt="Logo">
			<img src="/clown.png" alt="clown">
			</body></html>`,
			expected: []string{"https://crawler-test.com/logo.png", "https://crawler-test.com/clown.png"},
		},
		{
			name:     "Two other absolute",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			<img src="https://place-cat.org/a-cat.png" alt="Logo">
			<img src="https://place-dog.org/a-dog.png" alt="clown">
			</body></html>`,
			expected: []string{"https://place-cat.org/a-cat.png", "https://place-dog.org/a-dog.png"},
		},
		{
			name:      "Empty attribute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img alt="No source here!"><img src="/valid.png"></body></html>`,
			expected:  []string{"https://crawler-test.com/valid.png"},
		},
		{
			name:      "No images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body></body></html>`,
			expected:  []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
