package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(doc.Find("h1, h2").First().Text())
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	if main.Length() > 0 {
		if p := main.Find("p"); p.Length() > 0 {
			return strings.TrimSpace(p.First().Text())
		}
	}

	return strings.TrimSpace(doc.Find("p").First().Text())
}
