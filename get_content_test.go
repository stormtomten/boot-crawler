package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "First heading Basic",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "Second heading Basic",
			inputBody: "<html><body><h2>Test Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "No heading",
			inputBody: "<html><body><p>Test Title</p></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "Get First Paragraph main priority, one main paragraph",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "Get First Paragraph main priority, two main paragraph",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>First paragraph.</p>
					<p>Second paragraph.</p>
				</main>
			</body></html>`,
			expected: "First paragraph.",
		},
		{
			name: "No paragraph in main (fallback)",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<p>Second paragraph outside main</p>
				<main>
					<h1>Main paragraph.</h1>
				</main>
			</body></html>`,
			expected: "Outside paragraph.",
		},
		{
			name: "First paragraph empty (fallback)",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p></p>
					<p>Second paragraph.</p>
				</main>
			</body></html>`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}


