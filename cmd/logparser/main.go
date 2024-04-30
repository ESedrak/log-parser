package main

import (
	"log-parser/internal/log_mgmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logs, err := log_mgmt.LoadLogs()
	if err != nil {
		logrus.Error(err)
		return
	}

	logMatches, err := log_mgmt.ParseLogData(logs)
	if err != nil {
		logrus.Error(err)
		return
	}

	urlCounts, ipCounts := log_mgmt.CountLogMatchesIgnoresQuery(logMatches)

	uniqueIPs := ipCounts.UniqueIPs()

	//choose how many URLs or IP addresses to return (has to be greater than 0)
	requestedCountIP := 3
	requestedCountURL := 3

	mostActiveIPs, err := ipCounts.MostActiveIP(requestedCountIP)

	if err != nil {
		logrus.Error(err)
		return
	}

	topRequestedURLs, err := urlCounts.TopRequestedURLs(requestedCountURL)

	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("Unique IPs Count: ", uniqueIPs)
	logrus.Info("Most Active IPs: ", mostActiveIPs)
	logrus.Info("Top Requested URLs:", topRequestedURLs)
}
