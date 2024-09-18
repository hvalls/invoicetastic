package cmd

import (
	"invoicetastic/cmd/create_client"
	"invoicetastic/cmd/create_contact"
	"invoicetastic/cmd/create_invoice"
	"invoicetastic/cmd/create_paymentinfo"
	"invoicetastic/cmd/create_product"
	"invoicetastic/cmd/create_provider"
	"invoicetastic/cmd/create_tax"
	"invoicetastic/cmd/render"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoicestastic",
	Short: "Invoicestastic is a tool for generating invoices",
}

func init() {
	rootCmd.AddCommand(create_invoice.NewCmd())
	rootCmd.AddCommand(create_provider.NewCmd())
	rootCmd.AddCommand(create_client.NewCmd())
	rootCmd.AddCommand(create_contact.NewCmd())
	rootCmd.AddCommand(create_paymentinfo.NewCmd())
	rootCmd.AddCommand(create_product.NewCmd())
	rootCmd.AddCommand(create_tax.NewCmd())
	rootCmd.AddCommand(render.NewCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
