package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("Invalid url: %s: %s", str, err.Error())
	}

	switch u.Scheme {
	case "http":
		if u.Port() == "80" {
			u.Host = strings.TrimSuffix(u.Host, ":80")
		}

	case "https":
		if u.Port() == "443" {
			u.Host = strings.TrimSuffix(u.Host, ":443")
		}
	}

	u.Host = strings.TrimPrefix(u.Host, "www.")

	norm := strings.ToLower(fmt.Sprintf("%s%s", u.Host, u.Path))
	return norm, nil
}
