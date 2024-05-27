package main

import (
	"log-parser/internal/cli"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cli.New()

	cobra.CheckErr(rootCmd.Execute())
}
