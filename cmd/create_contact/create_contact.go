package create_contact

import (
	"fmt"
	"invoicetastic/contact"
	"invoicetastic/file"
	"invoicetastic/util"

	"github.com/spf13/cobra"
)

var name string
var email string
var website string

var createContactCmd = &cobra.Command{
	Use:   "create-contact",
	Short: "Create new contact YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		c := contact.New(name, email, website)
		filename := util.CleanString(name) + ".yml"
		err := file.WriteContent(filename, "Contact", c)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " contact created")
	},
}

func NewCmd() *cobra.Command {
	createContactCmd.Flags().StringVar(&name, "name", "", "Contact name")
	createContactCmd.MarkFlagRequired("name")
	createContactCmd.Flags().StringVar(&email, "email", "", "Contact email address")
	createContactCmd.Flags().StringVar(&website, "website", "", "Contact website")
	return createContactCmd
}
