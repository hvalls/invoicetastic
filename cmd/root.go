package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoicestastic",
	Short: "Invoicestastic is a tool for generating invoices",
}

func init() {
	rootCmd.AddCommand(buildGenerateCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
