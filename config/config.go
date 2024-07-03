package config

import (
	"invoicetastic/invoice"
	"os"

	"gopkg.in/yaml.v2"
)

const configFile = "invoicetastic.yml"

type Config struct {
	Defaults *invoice.Invoice `yaml:"defaults"`
}

func GetDefault() (*invoice.Invoice, error) {
	i, err := invoice.New()
	if err != nil {
		return nil, err
	}

	configFileContent, err := os.ReadFile(configFile)
	if err != nil {
		if err == os.ErrNotExist {
			return i, nil
		}
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal([]byte(configFileContent), &c)
	if err != nil {
		return nil, err
	}

	if c.Defaults.Provider != nil {
		i.Provider = c.Defaults.Provider
	}
	if c.Defaults.Customer != nil {
		i.Customer = c.Defaults.Customer
	}
	if len(c.Defaults.Products) > 0 {
		i.Products = c.Defaults.Products
	}
	if c.Defaults.Taxes != nil {
		i.Taxes = c.Defaults.Taxes
	}
	if c.Defaults.Contact != nil {
		i.Contact = c.Defaults.Contact
	}
	if c.Defaults.PaymentInfo != nil {
		i.PaymentInfo = c.Defaults.PaymentInfo
	}
	return i, nil
}
