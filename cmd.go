package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoicestastic",
	Short: "Invoicestastic is a tool for generating invoices",
}

var filePath string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoice",
	Run: func(cobraCmd *cobra.Command, args []string) {
		invoiceFileContent, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		inv, err := ParseInvoice(string(invoiceFileContent))
		if err != nil {
			panic(err)
		}

		templateFileContent, err := os.ReadFile(inv.Metadata.Template)
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New("latex").Parse(string(templateFileContent))
		if err != nil {
			panic(err)
		}

		texFileName := inv.Spec.Number + ".pdf"
		texFile, err := os.Create(texFileName)
		if err != nil {
			panic(err)
		}
		defer texFile.Close()

		err = tmpl.Execute(texFile, inv.Spec)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("pdflatex", texFileName)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Output:\n%s\n", output)
			panic(err)
		}

		err = os.Remove(inv.Spec.Number + ".aux")
		if err != nil {
			panic(err)
		}

		err = os.Remove(inv.Spec.Number + ".log")
		if err != nil {
			panic(err)
		}

		fmt.Println("✅" + inv.Spec.Number + ".pdf has been generated")
	},
}

func NewGenerateCmd() *cobra.Command {
	generateCmd.Flags().StringVarP(&filePath, "file", "f", "", "yaml file")
	return generateCmd
}
