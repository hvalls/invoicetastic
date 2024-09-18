package create_provider

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

var createProviderCmd = &cobra.Command{
	Use:   "create-provider",
	Short: "Create new provider YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		p := company.New(name, vatNumber, address)
		filename := util.CleanString(name) + ".yml"
		err := file.WriteContent(filename, "Provider", p)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " provider created")
	},
}

func NewCmd() *cobra.Command {
	createProviderCmd.Flags().StringVar(&name, "name", "", "Provider name")
	createProviderCmd.MarkFlagRequired("name")
	createProviderCmd.Flags().StringVar(&vatNumber, "vatnum", "", "Provider VAT number")
	createProviderCmd.Flags().StringArrayVar(&address, "address", []string{}, "Provider address line")
	return createProviderCmd
}
