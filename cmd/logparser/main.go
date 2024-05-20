package main

import (
	"log-parser/internal/cli"
	"log-parser/internal/config"
	"log-parser/internal/domain/log_mgmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func init() {
	config.Init("./config", "config", "json")
}

func main() {
	cfg := config.Values

	logChan := make(chan string)
	errChan := make(chan error)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(cfg.Path.LogPath, logChan, errChan)

	go log_mgmt.CountLogMatch(cfg.Regex.MatchIPsURLsIgnoreQuery, logChan, urlCountChan, ipCountChan)

	// receive errors and return if there were any errors loading or reading the file.
	err := <-errChan
	if err != nil {
		slog.Error("error", "err", err)
		return
	}

	// receive URL counts
	urlCount := <-urlCountChan

	// receive IP counts
	ipCount := <-ipCountChan

	rootCmd := cli.New(ipCount, urlCount)

	cobra.CheckErr(rootCmd.Execute())
}
