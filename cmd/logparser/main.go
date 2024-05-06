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
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(cfg.Path.FilePath, logChan)

	go log_mgmt.CountLogMatches(cfg.Regex.MatchIPsURlsIgnoreQuery, logChan, urlCountChan, ipCountChan)

	// receive URL counts
	urlCount := <-urlCountChan

	// receive IP counts
	ipCount := <-ipCountChan

	uniqueIPs := ip_mgmt.UniqueIPs(ipCount)

	mostActiveIPs, err := ip_mgmt.MostActiveIP(ipCount, cfg.Limit.MaxIPs)

	if err != nil {
		slog.Error("mostActiveIPs", "error", err)
		return
	}

	topRequestedURLs, err := url_mgmt.TopRequestedURLs(urlCount, cfg.Limit.MaxURLs)

	if err != nil {
		slog.Error("topRequestedURLs", "error", err)
		return
	}

	slog.Info("Unique IPs Count", "uniqueIPs", uniqueIPs)
	slog.Info("Most Active IPs: ", "activeIPs", mostActiveIPs)
	slog.Info("Top Requested URLs:", "requestedURLs", topRequestedURLs)
}
