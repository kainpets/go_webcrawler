package normalize_url

import (
	"strings"
	"net/url"
)

func NormalizeURL(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	host := u.Host

	if host == "" {
		host = u.Path
	}

	result := host + u.Path

	result = strings.TrimSuffix(result, "/")

	return result, nil
}