package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("invalid url: %s: %s", str, err.Error())
	}

	switch u.Scheme {
	case "http":
		if u.Port() == "80" {
			u.Host = u.Hostname()
		}

	case "https":
		if u.Port() == "443" {
			u.Host = u.Hostname()
		}
	}

	hostName := strings.TrimPrefix(u.Host, "www.")
	path := strings.TrimSuffix(u.Path, "/")

	norm := strings.ToLower(fmt.Sprintf("%s%s", hostName, path))
	return norm, nil
}
