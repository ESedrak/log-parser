package main

import (
	"log-parser/internal/config"
	"log-parser/internal/domain/ip_mgmt"
	"log-parser/internal/domain/log_mgmt"
	"log-parser/internal/domain/url_mgmt"
	"log/slog"
	"sync"
)

func init() {
	config.Init("config/config.json")
}

var wg sync.WaitGroup

func main() {
	cfg := config.Values

	wg.Add(2)

	logChan := make(chan string)
	errChan := make(chan error)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go func() {
		defer wg.Done()
		log_mgmt.LoadLogs(cfg.Path.LogPath, logChan, errChan)
	}()

	go func() {
		defer wg.Done()
		log_mgmt.CountLogMatches(cfg.Regex.MatchIPsURlsIgnoreQuery, logChan, urlCountChan, ipCountChan)
	}()

	// receive errors - if there was any error loading or reading the file. Return.
	err := <-errChan
	if err != nil {
		slog.Error("error", "err", err)
		return
	}

	// receive URL counts
	urlCount := <-urlCountChan

	// receive IP counts
	ipCount := <-ipCountChan

	wg.Wait()

	uniqueIPs, err := ip_mgmt.UniqueIPs(ipCount)

	if err != nil {
		slog.Error("uniqueIPs", "error", err)
		return
	}

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

	slog.Info("Unique IPs Count", "IPs", uniqueIPs)
	slog.Info("Most Active IPs: ", "IPs", mostActiveIPs)
	slog.Info("Top Requested URLs:", "URLs", topRequestedURLs)
}
