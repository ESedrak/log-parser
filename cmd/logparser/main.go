package main

import (
	"log-parser/internal/cli"
	"log-parser/internal/config"

	"github.com/spf13/cobra"
)

func init() {
	config.Init("./config", "config", "json")
}

func main() {
	rootCmd := cli.NewCli()

	cobra.CheckErr(rootCmd.Execute())
}
