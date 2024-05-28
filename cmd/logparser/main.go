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
	logPath := config.Values.Path.LogPath
	regex := config.Values.Regex.MatchIPsURLsIgnoreQuery

	rootCmd := cli.NewCli(logPath, regex)

	cobra.CheckErr(rootCmd.Execute())
}
