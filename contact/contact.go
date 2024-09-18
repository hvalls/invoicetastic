package contact

import (
	"invoicetastic/file"

	"gopkg.in/yaml.v2"
)

type Contact struct {
	Name    string `yaml:"name"`
	Email   string `yaml:"email"`
	Website string `yaml:"website"`
}

func LoadFrom(location string) (*Contact, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	var c Contact
	err = yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func New(name, email, website string) *Contact {
	return &Contact{
		Name:    name,
		Email:   email,
		Website: website,
	}
}
