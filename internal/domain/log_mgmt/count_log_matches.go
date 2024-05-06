package log_mgmt

import (
	"log/slog"
	"regexp"
)

/*
 * Function counts occurrences for the following HTTP methods: GET, PUT, DELETE, POST, HEAD.
 * Counts each unique IP address and URL(ignores queries)
 */
func CountLogMatches(regex string, logChan <-chan string, urlCountChan chan<- map[string]int, ipCountChan chan<- map[string]int) {
	// match the IP address and URL(ignore query). Used https://regex101.com/.
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
		} else {
			slog.Warn("log not parsed", "ignoring", log)
		}
	}

	urlCountChan <- urlCount
	ipCountChan <- ipCount

	close(urlCountChan)
	close(ipCountChan)
}
