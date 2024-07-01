package main

import (
	"gopkg.in/yaml.v2"
)

type TaxValue struct {
	Name       string `yaml:"name"`
	Percentage string `yaml:"percentage"`
	Value      string `yaml:"value"`
}

type Address struct {
	Line1 string `yaml:"line1"`
	Line2 string `yaml:"line2"`
	Line3 string `yaml:"line3"`
}

type Company struct {
	Name    string  `yaml:"name"`
	VAT     string  `yaml:"vat"`
	Address Address `yaml:"address"`
}

type Product struct {
	Description string `yaml:"description"`
	Qty         string `yaml:"qty"`
	UnitPrice   string `yaml:"unitPrice"`
	Total       string `yaml:"total"`
}

type Contact struct {
	Name    string `yaml:"name"`
	Email   string `yaml:"email"`
	Website string `yaml:"website"`
}

type PaymentInfo struct {
	Bank          string `yaml:"bank"`
	AccountName   string `yaml:"accountName"`
	AccountNumber string `yaml:"accountNumber"`
	SwiftBIC      string `yaml:"swiftBic"`
}

type Metadata struct {
	Template string `yaml:"template"`
}

type Spec struct {
	Number      string      `yaml:"number"`
	Date        string      `yaml:"date"`
	DueDate     string      `yaml:"dueDate"`
	Provider    Company     `yaml:"provider"`
	Customer    Company     `yaml:"customer"`
	Products    []Product   `yaml:"products"`
	Subtotal    string      `yaml:"subtotal"`
	Taxes       []TaxValue  `yaml:"taxes"`
	Total       string      `yaml:"total"`
	Contact     Contact     `yaml:"contact"`
	PaymentInfo PaymentInfo `yaml:"paymentInfo"`
}

type Invoice struct {
	Metadata Metadata `yaml:"metadata"`
	Spec     Spec     `yaml:"spec"`
}

func ParseInvoice(content string) (*Invoice, error) {
	var inv Invoice
	err := yaml.Unmarshal([]byte(content), &inv)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}
