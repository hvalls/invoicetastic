package cmd

import (
	"fmt"
	"invoicetastic/invoice"
	"invoicetastic/latextemplate"

	"github.com/spf13/cobra"
)

var invoiceLocation string
var templateLocation string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoice PDF",
	Run: func(cobraCmd *cobra.Command, args []string) {
		inv, err := invoice.NewFrom(invoiceLocation)
		if err != nil {
			panic(err)
		}

		t, err := latextemplate.New(templateLocation)
		if err != nil {
			panic(err)
		}

		fileName, err := t.RenderPDF(inv.Number, inv)
		if err != nil {
			panic(err)
		}

		fmt.Println(fileName + " file created")
	},
}

func buildGenerateCmd() *cobra.Command {
	generateCmd.Flags().StringVarP(&invoiceLocation, "file", "f", "", "invoice yaml file location (file path or URL)")
	err := generateCmd.MarkFlagRequired("file")
	if err != nil {
		panic(err)
	}

	generateCmd.Flags().StringVarP(&templateLocation, "template", "t", "", "template location (file path or URL)")

	return generateCmd
}
