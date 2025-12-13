package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and www.",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and www. and query",
			inputURL: "https://www.blog.boot.dev/path/search?q=none",
			expected: "blog.boot.dev/path/search",
		},
		{
			name:     "remove scheme and www. and empty query",
			inputURL: "https://www.blog.boot.dev/path/search?",
			expected: "blog.boot.dev/path/search",
		},
		{
			name:     "remove scheme in capitals",
			inputURL: "HTTP://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "wierd but legal path",
			inputURL: "HTTP://blog.boot.dev//path//to///page",
			expected: "blog.boot.dev//path//to///page",
		},
		{
			name:     "wierd but legal path",
			inputURL: "HTTP://blog.boot.dev//path//..//to///page",
			expected: "blog.boot.dev//path//..//to///page",
		},
		{
			name:     "Uppercase",
			inputURL: "HTTP://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "Explicit ports http remove",
			inputURL: "http://blog.boot.dev:80/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "Explicit ports https keep",
			inputURL: "https://blog.boot.dev:80/path",
			expected: "blog.boot.dev:80/path",
		},
		{
			name:     "Explicit ports http keep",
			inputURL: "http://blog.boot.dev:443/path",
			expected: "blog.boot.dev:443/path",
		},
		{
			name:     "Explicit ports https remove",
			inputURL: "https://blog.boot.dev:443/path",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
