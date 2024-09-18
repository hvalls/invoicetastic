package create_product

import (
	"fmt"
	"invoicetastic/file"
	"invoicetastic/product"
	"invoicetastic/util"

	"github.com/spf13/cobra"
)

var name string
var unitPrice float64

var createProductCmd = &cobra.Command{
	Use:   "create-product",
	Short: "Create new product YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		p := product.New(name, unitPrice)
		filename := util.CleanString(name) + ".yml"
		err := file.WriteContent(filename, "Product", p)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " product created")
	},
}

func NewCmd() *cobra.Command {
	createProductCmd.Flags().StringVar(&name, "name", "", "Product name")
	createProductCmd.MarkFlagRequired("name")
	createProductCmd.Flags().Float64Var(&unitPrice, "unitprice", 0, "Product unit price")
	return createProductCmd
}
