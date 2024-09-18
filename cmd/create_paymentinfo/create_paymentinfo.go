package create_paymentinfo

import (
	"fmt"
	"invoicetastic/file"
	"invoicetastic/paymentinfo"
	"invoicetastic/util"

	"github.com/spf13/cobra"
)

var bank string
var accountName string
var accountNumber string
var swiftBIC string

var createPaymentInfo = &cobra.Command{
	Use:   "create-paymentinfo",
	Short: "Create new provider YAML file",
	Run: func(cobraCmd *cobra.Command, args []string) {
		p := paymentinfo.New(bank, accountName, accountNumber, swiftBIC)
		filename := util.CleanString(bank) + ".yml"
		err := file.WriteContent(filename, "PaymentInfo", p)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename + " paymentinfo created")
	},
}

func NewCmd() *cobra.Command {
	createPaymentInfo.Flags().StringVar(&bank, "bank", "", "Bank name")
	createPaymentInfo.MarkFlagRequired("bank")
	createPaymentInfo.Flags().StringVar(&accountName, "accountname", "", "Bank account name")
	createPaymentInfo.Flags().StringVar(&accountNumber, "accountnum", "", "Bank account number")
	createPaymentInfo.Flags().StringVar(&swiftBIC, "swiftbic", "", "Bank account Swift/BIC")
	return createPaymentInfo
}
