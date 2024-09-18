package create_tax

import (
	"fmt"
	"invoicetastic/file"
	"invoicetastic/tax"
	"invoicetastic/util"

	"github.com/spf13/cobra"
)

var name string
var percentage float64

var createTaxCmd = &cobra.Command{
	Use:   "create-tax",
	Short: "Create new tax YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		t := tax.New(name, percentage)
		filename := util.CleanString(name) + ".yml"
		err := file.WriteContent(filename, "Tax", t)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " tax created")
	},
}

func NewCmd() *cobra.Command {
	createTaxCmd.Flags().StringVar(&name, "name", "", "Tax name")
	createTaxCmd.MarkFlagRequired("name")
	createTaxCmd.Flags().Float64Var(&percentage, "percentage", 0, "Tax percentage")
	return createTaxCmd
}
