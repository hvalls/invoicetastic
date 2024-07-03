package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [invoice number]",
	Short: "Create new invoice YAML file",
	Args:  cobra.ExactArgs(1),
	Run: func(cobraCmd *cobra.Command, args []string) {
		invoiceNumber := args[0]
		text := `number: "` + invoiceNumber + `"
date: ""
dueDate: ""
provider:
  name: ""
  vat: ""
  address: 
    line1: ""
    line2: "" 
    line3: ""
customer:
  name: ""
  vat: ""
  address:
    line1: ""
    line2: ""
    line3: ""
products:
  - description: ""
    qty: 0
    unitPrice: 0.0
taxes:
  - name: ""
    percentage: 0
contact:
  name: ""
  email: ""
  website: ""
paymentInfo:
  bank: ""
  accountName: ""
  accountNumber: ""
  swiftBic: ""
`
		filename := invoiceNumber + ".yml"

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		_, err = file.Write([]byte(text))
		if err != nil {
			panic(err)
		}

		err = file.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println(invoiceNumber + ".yml file created")
	},
}

func buildCreateCmd() *cobra.Command {
	return createCmd
}
