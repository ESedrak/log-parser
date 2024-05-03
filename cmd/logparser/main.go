package main

import (
	"log-parser/internal/ip_mgmt"
	"log-parser/internal/log_mgmt"
	"log-parser/internal/url_mgmt"
	"log/slog"
)

func main() {
	filename := "logs/log_file.log"

	logChan := make(chan string)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(filename, logChan)

	go log_mgmt.CountLogMatchesIgnoresQuery(logChan, urlCountChan, ipCountChan)

	// receive URL counts
	urlCount := <-urlCountChan

	// receive IP counts
	ipCount := <-ipCountChan

	uniqueIPs := ip_mgmt.UniqueIPs(ipCount)

	//choose how many URLs or IP addresses to return (has to be greater than 0)
	requestedNumIP := 3
	requestedNumURL := 3

	mostActiveIPs, err := ip_mgmt.MostActiveIP(ipCount, requestedNumIP)

	if err != nil {
		slog.Error("mostActiveIPs", "error", err)
		return
	}

	topRequestedURLs, err := url_mgmt.TopRequestedURLs(urlCount, requestedNumURL)

	if err != nil {
		slog.Error("topRequestedURLs", "error", err)
		return
	}

	slog.Info("Unique IPs Count", "uniqueIPs", uniqueIPs)
	slog.Info("Most Active IPs: ", "activeIPs", mostActiveIPs)
	slog.Info("Top Requested URLs:", "requestedURLs", topRequestedURLs)
}
