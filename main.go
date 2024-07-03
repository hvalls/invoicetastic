package main

import "invoicetastic/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
