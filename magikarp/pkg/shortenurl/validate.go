package shortenurl

import (
	"net/url"
)

func validateURL(inputURL string) error {
	_, err := url.ParseRequestURI(inputURL)
	return err
}
