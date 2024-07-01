package main

func main() {
	rootCmd.AddCommand(NewGenerateCmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
