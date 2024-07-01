package main

import (
	"gopkg.in/yaml.v2"
)

type TaxValue struct {
	Name       string  `yaml:"name"`
	Percentage float64 `yaml:"percentage"`
	Value      float64 // Computed
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
	Description string  `yaml:"description"`
	Qty         float64 `yaml:"qty"`
	UnitPrice   float64 `yaml:"unitPrice"`
	Total       float64 // Computed
}

func (p *Product) getTotal() float64 {
	return p.UnitPrice * p.Qty
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

type Invoice struct {
	Number      string      `yaml:"number"`
	Date        string      `yaml:"date"`
	DueDate     string      `yaml:"dueDate"`
	Provider    Company     `yaml:"provider"`
	Customer    Company     `yaml:"customer"`
	Products    []*Product  `yaml:"products"`
	Subtotal    float64     // Computed
	Taxes       []*TaxValue `yaml:"taxes"`
	Total       float64     // Computed
	Contact     Contact     `yaml:"contact"`
	PaymentInfo PaymentInfo `yaml:"paymentInfo"`
}

func ParseInvoice(content string) (*Invoice, error) {
	var inv Invoice
	err := yaml.Unmarshal([]byte(content), &inv)
	if err != nil {
		return nil, err
	}

	// compute product totals
	for _, p := range inv.Products {
		p.Total = p.getTotal()
	}

	// compute invoice subtotal
	subtotal, err := inv.getSubtotal()
	if err != nil {
		return nil, err
	}
	inv.Subtotal = subtotal

	// compute tax values
	for _, t := range inv.Taxes {
		t.Value = (inv.Subtotal * t.Percentage) / 100
	}

	// compute invoice total
	inv.Total = inv.getTotal()

	return &inv, nil
}

func (inv *Invoice) getSubtotal() (float64, error) {
	subtotal := 0.0
	for _, p := range inv.Products {
		subtotal += p.getTotal()
	}
	return subtotal, nil
}

func (inv *Invoice) getTotal() float64 {
	total := inv.Subtotal
	for _, tv := range inv.Taxes {
		total += tv.Value
	}
	return total
}
