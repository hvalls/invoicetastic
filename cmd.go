package main

import (
	"fmt"
	"invoicetastic/latextemplate"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoicestastic",
	Short: "Invoicestastic is a tool for generating invoices",
}

var filePath string
var templatePath string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoice",
	Run: func(cobraCmd *cobra.Command, args []string) {
		invoiceFileContent, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		inv, err := ParseInvoice(string(invoiceFileContent))
		if err != nil {
			panic(err)
		}

		t, err := latextemplate.New(templatePath)
		if err != nil {
			panic(err)
		}

		err = t.Render(inv.Number, inv)
		if err != nil {
			panic(err)
		}

		fmt.Println("✅ Invoice has been generated")
	},
}

func NewGenerateCmd() *cobra.Command {
	generateCmd.Flags().StringVarP(&filePath, "file", "f", "", "yaml file")
	err := generateCmd.MarkFlagRequired("file")
	if err != nil {
		panic(err)
	}

	generateCmd.Flags().StringVarP(&templatePath, "template", "t", "", "yaml file")

	return generateCmd
}
