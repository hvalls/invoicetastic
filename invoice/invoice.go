package invoice

import (
	"invoicetastic/company"
	"invoicetastic/contact"
	"invoicetastic/file"
	"invoicetastic/paymentinfo"
	"invoicetastic/product"
	"invoicetastic/tax"
	"time"

	"gopkg.in/yaml.v2"
)

type Invoice struct {
	Number      string                   `yaml:"number"`
	Date        string                   `yaml:"date"`
	DueDate     string                   `yaml:"dueDate"`
	Provider    *company.Company         `yaml:"provider"`
	Client      *company.Company         `yaml:"client"`
	Lines       []*Line                  `yaml:"lines"`
	Taxes       []*tax.Tax               `yaml:"taxes"`
	Contact     *contact.Contact         `yaml:"contact"`
	PaymentInfo *paymentinfo.PaymentInfo `yaml:"paymentInfo"`
	Subtotal    float64                  `yaml:"-"` // Computed
	Total       float64                  `yaml:"-"` // Computed
}

func New() *Invoice {
	return &Invoice{
		Number:  "",
		Date:    time.Now().Format("2006-01-02"),
		DueDate: time.Now().Add(24 * time.Hour * 30).Format("2006-01-02"),
		Provider: &company.Company{
			Address: []string{},
		},
		Client: &company.Company{
			Address: []string{},
		},
		Lines:       []*Line{},
		Subtotal:    0,
		Taxes:       []*tax.Tax{},
		Total:       0,
		Contact:     &contact.Contact{},
		PaymentInfo: &paymentinfo.PaymentInfo{},
	}
}

func (i *Invoice) AddLine(p *product.Product, qty float64) {
	i.Lines = append(i.Lines, &Line{
		Name:      p.Name,
		UnitPrice: p.UnitPrice,
		Qty:       qty,
	})
}

func (i *Invoice) AddTax(t *tax.Tax) {
	i.Taxes = append(i.Taxes, t)
}

func NewFrom(location string) (*Invoice, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	return newFromContent(content)
}

func newFromContent(content string) (*Invoice, error) {
	var inv Invoice
	err := yaml.Unmarshal([]byte(content), &inv)
	if err != nil {
		return nil, err
	}

	// compute lines totals
	for _, l := range inv.Lines {
		l.Total = l.getTotal()
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
	for _, p := range inv.Lines {
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
