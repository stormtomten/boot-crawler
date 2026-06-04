package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "Basic",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:     "Two links",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			<a href="https://crawler-test.com"><span>Boot.dev</span></a>
			<a href="https://crawler-test.com"><span>Boot.dev</span></a>
			</body></html>`,
			expected: []string{"https://crawler-test.com", "https://crawler-test.com"},
		},
		{
			name:     "No link",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			</body></html>`,
			expected: []string{},
		},
		{
			name:      "Relative /about",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/about"><span>About</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/about"},
		},
		{
			name:      "Relative /about/hurr",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/about/hurr"><span>Hurr?</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/about/hurr"},
		},
		{
			name:      "Some other place",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://some-other.place"><span>That one other place</span></a></body></html>`,
			expected:  []string{"https://some-other.place"},
		},
		{
			name:     "Empty atribute",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			<a><span>No Source</span></a>
			<a href="/minimi"><span>Minimi</span></a>
			</body></html>`,
			expected: []string{"https://crawler-test.com/minimi"},
		},
		{
			name:      "Nested whitespace",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href=" /relative-path "><span>Click Me</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/relative-path"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
