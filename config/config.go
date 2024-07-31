package config

import (
	"errors"
	"invoicetastic/invoice"
	"os"

	"gopkg.in/yaml.v2"
)

const configFile = "invoicetastic.yml"

type Defaults struct {
	Template string           `yaml:"template"`
	Data     *invoice.Invoice `yaml:"data"`
}

type Config struct {
	Defaults *Defaults `yaml:"defaults"`
}

func GetDefault() (*invoice.Invoice, error) {
	i, err := invoice.New()
	if err != nil {
		return nil, err
	}

	configFileContent, err := os.ReadFile(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return i, nil
		}
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal([]byte(configFileContent), &c)
	if err != nil {
		return nil, err
	}

	if c.Defaults.Data.Provider != nil {
		i.Provider = c.Defaults.Data.Provider
	}
	if c.Defaults.Data.Customer != nil {
		i.Customer = c.Defaults.Data.Customer
	}
	if len(c.Defaults.Data.Products) > 0 {
		i.Products = c.Defaults.Data.Products
	}
	if c.Defaults.Data.Taxes != nil {
		i.Taxes = c.Defaults.Data.Taxes
	}
	if c.Defaults.Data.Contact != nil {
		i.Contact = c.Defaults.Data.Contact
	}
	if c.Defaults.Data.PaymentInfo != nil {
		i.PaymentInfo = c.Defaults.Data.PaymentInfo
	}
	return i, nil
}
