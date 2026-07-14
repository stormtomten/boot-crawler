package main

import "net/url"

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	resUrl, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	links, err := getURLsFromHTML(html, resUrl)
	if err != nil {
		links = nil
	}
	images, err := getImagesFromHTML(html, resUrl)
	if err != nil {
		images = nil
	}

	return PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  links,
		ImageURLs:      images,
	}
}
