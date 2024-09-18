package company

import (
	"invoicetastic/file"

	"gopkg.in/yaml.v2"
)

type Company struct {
	Name      string   `yaml:"name"`
	VATNumber string   `yaml:"vatNumber"`
	Address   []string `yaml:"address"`
}

func LoadFrom(location string) (*Company, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	var c Company
	err = yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func New(name, vatNumber string, address []string) *Company {
	return &Company{
		Name:      name,
		VATNumber: vatNumber,
		Address:   address,
	}
}
