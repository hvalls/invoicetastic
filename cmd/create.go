package cmd

import (
	"fmt"
	"invoicetastic/config"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [invoice number]",
	Short: "Create new invoice YAML file",
	Args:  cobra.ExactArgs(1),
	Run: func(cobraCmd *cobra.Command, args []string) {
		i, err := config.GetDefault()
		if err != nil {
			panic(err)
		}
		invoiceNumber := args[0]
		i.Number = invoiceNumber

		filename := invoiceNumber + ".yml"
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		text, err := i.MarshalYAML()
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
