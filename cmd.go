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

		templateFileContent, err := os.ReadFile("_templates/spanish-default.tex")
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New("latex").Parse(string(templateFileContent))
		if err != nil {
			panic(err)
		}

		texFileName := inv.Number + ".pdf"
		texFile, err := os.Create(texFileName)
		if err != nil {
			panic(err)
		}
		defer texFile.Close()

		err = tmpl.Execute(texFile, inv)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("pdflatex", texFileName)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Output:\n%s\n", output)
			panic(err)
		}

		err = os.Remove(inv.Number + ".aux")
		if err != nil {
			panic(err)
		}

		err = os.Remove(inv.Number + ".log")
		if err != nil {
			panic(err)
		}

		fmt.Println("✅" + inv.Number + ".pdf has been generated")
	},
}

func NewGenerateCmd() *cobra.Command {
	generateCmd.Flags().StringVarP(&filePath, "file", "f", "", "yaml file")
	return generateCmd
}
