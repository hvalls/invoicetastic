package util

import "net/url"

func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
