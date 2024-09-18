package product

import (
	"invoicetastic/file"

	"gopkg.in/yaml.v2"
)

type Product struct {
	Name      string  `yaml:"name"`
	UnitPrice float64 `yaml:"unitPrice"`
}

func LoadFrom(location string) (*Product, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	var p Product
	err = yaml.Unmarshal([]byte(content), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func New(name string, unitPrice float64) *Product {
	return &Product{
		Name:      name,
		UnitPrice: unitPrice,
	}
}
