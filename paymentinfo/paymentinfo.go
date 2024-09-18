package paymentinfo

import (
	"invoicetastic/file"

	"gopkg.in/yaml.v2"
)

type PaymentInfo struct {
	Bank          string `yaml:"bank"`
	AccountName   string `yaml:"accountName"`
	AccountNumber string `yaml:"accountNumber"`
	SwiftBIC      string `yaml:"swiftBic"`
}

func LoadFrom(location string) (*PaymentInfo, error) {
	content, err := file.ReadContent(location)
	if err != nil {
		return nil, err
	}
	var p PaymentInfo
	err = yaml.Unmarshal([]byte(content), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func New(bank, accountName, accountNumber, swiftBIC string) *PaymentInfo {
	return &PaymentInfo{
		Bank:          bank,
		AccountName:   accountName,
		AccountNumber: accountNumber,
		SwiftBIC:      swiftBIC,
	}
}
