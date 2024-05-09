package main

import (
	"log-parser/internal/config"
	"log-parser/internal/domain/ip_mgmt"
	"log-parser/internal/domain/log_mgmt"
	"log-parser/internal/domain/url_mgmt"
	"log/slog"
)

func init() {
	config.Init("config/config.json")
}

func main() {
	cfg := config.Values

	logChan := make(chan string)
	errChan := make(chan error)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(cfg.Path.LogPath, logChan, errChan)

	go log_mgmt.CountLogMatch(cfg.Regex.MatchIPsURlsIgnoreQuery, logChan, urlCountChan, ipCountChan)

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

	uniqueIPs, err := ip_mgmt.UniqueIPs(ipCount)

	if err != nil {
		slog.Error("uniqueIPs", "error", err)
		return
	}

	mostActiveIPs, err := ip_mgmt.MostActiveIP(ipCount, cfg.RequestedNum.IP)

	if err != nil {
		slog.Error("mostActiveIPs", "error", err)
		return
	}

	topRequestedURLs, err := url_mgmt.TopRequestedURLs(urlCount, cfg.RequestedNum.URL)

	if err != nil {
		slog.Error("topRequestedURLs", "error", err)
		return
	}

	slog.Info("Unique IPs Count", "IPs", uniqueIPs)
	slog.Info("Most Active IPs: ", "IPs", mostActiveIPs)
	slog.Info("Top Requested URLs:", "URLs", topRequestedURLs)
}
