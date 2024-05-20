package log_mgmt

import (
	"log/slog"
	"regexp"
)

/*
 * Count each log match based on the given regex (used https://regex101.com/)
 * Default regex: captures IP and URL(ignores any query) for HTTP methods: GET, PUT, DELETE, POST, HEAD
 */
func CountLogMatch(regex string, logChan <-chan string, urlCountChan chan<- map[string]int, ipCountChan chan<- map[string]int) {
	defer close(urlCountChan)
	defer close(ipCountChan)

	logRegex := regexp.MustCompile(regex)

	urlCount := make(map[string]int)
	ipCount := make(map[string]int)

	for log := range logChan {
		// find matches in the log
		matches := logRegex.FindStringSubmatch(log)

		// if successful - will have 3 matches (full log and the two capture groups: ip and url)
		if len(matches) == 3 {
			ip := matches[1]
			url := matches[2]

			urlCount[url]++
			ipCount[ip]++

			slog.Info("log parsed successfully", "IP", ip, "URL", url)
		} else {
			slog.Warn("log not parsed", "ignoring", log)
		}
	}

	urlCountChan <- urlCount
	ipCountChan <- ipCount
}
