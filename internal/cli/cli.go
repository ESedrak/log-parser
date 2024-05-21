package cli

import (
	"log-parser/internal/cli/ipcli"
	"log-parser/internal/cli/urlcli"
	"log-parser/internal/config"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

func init() {
	config.Init("./config", "config", "json")
}

func New(ipCount, urlCount map[string]int) *cobra.Command {

	var logparser = &cobra.Command{
		Use:   "logparser",
		Short: "LogParser CLI application",
		Long:  "Logparser CLI application\nUse this CLI for all your parsing needs.",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	useDefault := askUseDefaultValues()
		// 	cfg := config.Values

		// 	requestedNumIP := cfg.RequestedNum.IP
		// 	requestedNumURL := cfg.RequestedNum.URL

		// 	if !useDefault {
		// 		requestedNumIP = getRequestCount("IPs", requestedNumIP)
		// 		requestedNumURL = getRequestCount("URLs", requestedNumURL)
		// 	}

		// 	uniqueIPs, err := ip_mgmt.UniqueIPs(ipCount)

		// 	if err != nil {
		// 		slog.Error("uniqueIPs", "error", err)
		// 		return
		// 	}

		// 	mostActiveIPs, err := ip_mgmt.MostActiveIP(ipCount, requestedNumIP)

		// 	if err != nil {
		// 		slog.Error("mostActiveIPs", "error", err)
		// 		return
		// 	}

		// 	topRequestedURLs, err := url_mgmt.TopRequestedURLs(urlCount, requestedNumURL)

		// 	if err != nil {
		// 		slog.Error("topRequestedURLs", "error", err)
		// 		return
		// 	}

		// 	slog.Info("Unique IPs Count", "IPs", uniqueIPs)
		// 	slog.Info("Most Active IPs: ", "IPs", mostActiveIPs)
		// 	slog.Info("Top Requested URLs:", "URLs", topRequestedURLs)
		// },
	}

	logparser.AddCommand(ipcli.NewIP(ipCount), urlcli.NewURL(urlCount))

	return logparser
}

func getRequestCount(target string, defaultValue int) int {
	// Read the default value from the config
	prompt := promptui.Prompt{
		Label:   "Enter the number to return for: " + target,
		Default: strconv.Itoa(defaultValue),
	}

	result, err := prompt.Run()
	if err != nil {
		slog.Error("prompt failed", "err", err)
		os.Exit(1)
	}

	requestedCount, err := strconv.Atoi(result)
	if err != nil {
		slog.Error("Invalid input", "err", err)
		os.Exit(1)
	}

	return requestedCount
}

func askUseDefaultValues() bool {
	prompt := promptui.Prompt{
		Label: "Use default values for IP and URL counts? (y/n)",
	}

	result, err := prompt.Run()
	if err != nil {
		slog.Error("prompt failed", "err", err)
		os.Exit(1)
	}

	// If the result is 'y' or 'Y', use the default values
	return result == "y" || result == "Y"
}
