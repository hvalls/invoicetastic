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
var templateLocation string

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

		t, err := latextemplate.New(templateLocation)
		if err != nil {
			panic(err)
		}

		err = t.RenderPDF(inv.Number, inv)
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

	generateCmd.Flags().StringVarP(&templateLocation, "template", "t", "", "template location (file path or URL)")

	return generateCmd
}
