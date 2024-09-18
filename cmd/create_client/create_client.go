package create_client

import (
	"fmt"
	"invoicetastic/company"
	"invoicetastic/file"
	"invoicetastic/util"

	"github.com/spf13/cobra"
)

var name string
var vatNumber string
var address []string

var createClientCmd = &cobra.Command{
	Use:   "create-client",
	Short: "Create new client YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		c := company.New(name, vatNumber, address)
		filename := util.CleanString(name) + ".yml"
		err := file.WriteContent(filename, "Client", c)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " client created")
	},
}

func NewCmd() *cobra.Command {
	createClientCmd.Flags().StringVar(&name, "name", "", "Provider name")
	err := createClientCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	createClientCmd.Flags().StringVar(&vatNumber, "vatnum", "", "Provider VAT number")
	createClientCmd.Flags().StringArrayVar(&address, "address", []string{}, "Provider address line")
	return createClientCmd
}
