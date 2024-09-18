package render

import (
	"fmt"
	"invoicetastic/invoice"
	"invoicetastic/latextemplate"
	"os"

	"github.com/spf13/cobra"
)

var invoiceLocation string
var templateLocation string

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render invoice PDF",
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
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(fileName + " file created")
	},
}

func NewCmd() *cobra.Command {
	renderCmd.Flags().StringVar(&invoiceLocation, "file", "", "invoice yaml file location (file path or URL)")
	renderCmd.MarkFlagRequired("file")
	renderCmd.Flags().StringVar(&templateLocation, "template", "", "template location (file path or URL)")
	renderCmd.MarkFlagRequired("template")
	return renderCmd
}
