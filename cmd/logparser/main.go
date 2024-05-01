package main

import (
	"log-parser/internal/log_mgmt"
	"log/slog"
)

func main() {
	logs, err := log_mgmt.LoadLogs()
	if err != nil {
		slog.Error("unable to load logs", "error", err)
		return
	}

	logMatches, err := log_mgmt.ParseLogData(logs)
	if err != nil {
		slog.Error("unable to parse logs", "error", err)
		return
	}

	urlCounts, ipCounts := log_mgmt.CountLogMatchesIgnoresQuery(logMatches)

	uniqueIPs := ipCounts.UniqueIPs()

	//choose how many URLs or IP addresses to return (has to be greater than 0)
	requestedCountIP := 3
	requestedCountURL := 3

	mostActiveIPs, err := ipCounts.MostActiveIP(requestedCountIP)

	if err != nil {
		slog.Error("mostActiveIPs", "error", err)
		return
	}

	topRequestedURLs, err := urlCounts.TopRequestedURLs(requestedCountURL)

	if err != nil {
		slog.Error("topRequestedURLs", "error", err)
		return
	}

	slog.Info("Unique IPs Count", "uniqueIPs", uniqueIPs)
	slog.Info("Most Active IPs: ", "activeIPs", mostActiveIPs)
	slog.Info("Top Requested URLs:", "requestedURLs", topRequestedURLs)
}
