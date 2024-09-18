package tax

import (
	"invoicetastic/file"

	"gopkg.in/yaml.v2"
)

type Tax struct {
	Name       string  `yaml:"name"`
	Percentage float64 `yaml:"percentage"`
	Value      float64 `yaml:"-"` // Computed
}

func LoadFrom(location string) (*Tax, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	var t Tax
	err = yaml.Unmarshal([]byte(content), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func New(name string, percentage float64) *Tax {
	return &Tax{
		Name:       name,
		Percentage: percentage,
	}
}
