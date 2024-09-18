package create_invoice

import (
	"fmt"
	"invoicetastic/company"
	"invoicetastic/contact"
	"invoicetastic/file"
	"invoicetastic/invoice"
	"invoicetastic/paymentinfo"
	"invoicetastic/product"
	"invoicetastic/tax"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var date string
var dueDate string
var providerLocation string
var clientLocation string
var lines []string
var taxes []string
var contactLocation string
var paymentInfoLocation string

var createInvoiceCmd = &cobra.Command{
	Use:   "create-invoice [invoice number]",
	Short: "Create new invoice YAML file",
	Args:  cobra.ExactArgs(1),
	Run: func(cobraCmd *cobra.Command, args []string) {
		i := invoice.New()
		i.Number = args[0]

		if date != "" {
			i.Date = date
		}
		if dueDate != "" {
			i.DueDate = dueDate
		}

		i.Provider = company.New("", "", []string{})
		if providerLocation != "" {
			prov, err := company.LoadFrom(providerLocation)
			if err != nil {
				panic(err)
			}
			i.Provider = prov
		}

		i.Client = company.New("", "", []string{})
		if clientLocation != "" {
			client, err := company.LoadFrom(clientLocation)
			if err != nil {
				panic(err)
			}
			i.Client = client
		}

		i.Contact = contact.New("", "", "")
		if contactLocation != "" {
			cont, err := contact.LoadFrom(contactLocation)
			if err != nil {
				panic(err)
			}
			i.Contact = cont
		}

		i.PaymentInfo = paymentinfo.New("", "", "", "")
		if paymentInfoLocation != "" {
			pInfo, err := paymentinfo.LoadFrom(paymentInfoLocation)
			if err != nil {
				panic(err)
			}
			i.PaymentInfo = pInfo
		}

		for _, line := range lines {
			parts := strings.Split(line, ":")
			prod, err := product.LoadFrom(parts[0])
			if err != nil {
				panic(err)
			}
			qty, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				panic(err)
			}
			i.AddLine(prod, qty)
		}
		for _, t := range taxes {
			tax, err := tax.LoadFrom(t)
			if err != nil {
				panic(err)
			}
			i.AddTax(tax)
		}

		filename := i.Number + ".yml"
		err := file.WriteContent(filename, "Invoice", i)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " invoice created")
	},
}

func NewCmd() *cobra.Command {
	createInvoiceCmd.Flags().StringVar(&date, "date", "", "Invoice date")
	createInvoiceCmd.Flags().StringVar(&dueDate, "duedate", "", "Invoice due date")
	createInvoiceCmd.Flags().StringVar(&providerLocation, "provider", "", "Provider YAML path")
	createInvoiceCmd.Flags().StringVar(&clientLocation, "client", "", "Client YAML path")
	createInvoiceCmd.Flags().StringArrayVar(&lines, "line", []string{}, "Invoice line composed of a product YAML path and quantity, separated by colon. e.g. product1.yml:4")
	createInvoiceCmd.Flags().StringArrayVar(&taxes, "tax", []string{}, "Tax YAML path")
	createInvoiceCmd.Flags().StringVar(&contactLocation, "contact", "", "Contact YAML path")
	createInvoiceCmd.Flags().StringVar(&paymentInfoLocation, "payment", "", "PaymentInfo YAML path")
	return createInvoiceCmd
}
