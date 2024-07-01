package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Tax struct {
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
	Import  string  `yaml:"import"`
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
	Provider    *Company    `yaml:"provider"`
	Customer    *Company    `yaml:"customer"`
	Products    []*Product  `yaml:"products"`
	Subtotal    float64     // Computed
	Taxes       any         `yaml:"taxes"`
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

	// import provider
	if inv.Provider.Import != "" {
		prov, err := importCompany(inv.Provider.Import)
		if err != nil {
			return nil, err
		}
		inv.Provider = prov
	}

	// import customer
	if inv.Customer.Import != "" {
		cust, err := importCompany(inv.Customer.Import)
		if err != nil {
			return nil, err
		}
		inv.Customer = cust
	}

	// import taxes
	if _, ok := inv.Taxes.([]any); !ok {
		im := inv.Taxes.(map[any]any)
		txx, err := importTaxes(im["import"].(string))
		if err != nil {
			return nil, err
		}
		inv.Taxes = txx
	} else {
		txx, err := toTaxSlice(inv.Taxes.([]any))
		if err != nil {
			return nil, err
		}
		inv.Taxes = txx
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
	for _, t := range inv.Taxes.([]*Tax) {
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
	for _, tv := range inv.Taxes.([]*Tax) {
		total += tv.Value
	}
	return total
}

func importCompany(filePath string) (*Company, error) {
	companyFileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var comp Company
	err = yaml.Unmarshal([]byte(companyFileContent), &comp)
	if err != nil {
		return nil, err
	}
	return &comp, nil
}

func importTaxes(filePath string) ([]*Tax, error) {
	taxesFileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	txx := make([]*Tax, 0)
	err = yaml.Unmarshal([]byte(taxesFileContent), &txx)
	if err != nil {
		return nil, err
	}
	return txx, nil
}

func toTaxSlice(items []any) ([]*Tax, error) {
	txx := make([]*Tax, 0)
	for _, item := range items {
		taxMap, ok := item.(map[any]any)
		if !ok {
			return nil, fmt.Errorf("unexpected tax format")
		}
		name, ok := taxMap["name"].(string)
		if !ok {
			return nil, fmt.Errorf("tax name is not a string")
		}
		percentage, err := convertToFloat64(taxMap["percentage"])
		if err != nil {
			return nil, err
		}
		tax := &Tax{
			Name:       name,
			Percentage: percentage,
		}
		txx = append(txx, tax)
	}
	return txx, nil
}

func convertToFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse float from string: %v", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unexpected type %T for float64 conversion", val)
	}
}
