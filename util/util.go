package util

import (
	"net/url"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func CleanString(name string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := re.ReplaceAllString(name, "_")
	cleaned = strings.ToLower(cleaned)
	cleaned = strings.Trim(cleaned, "_")
	return cleaned
}

func MarshalYAML(o any) (string, error) {
	bb, err := yaml.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}
