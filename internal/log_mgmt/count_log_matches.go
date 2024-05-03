package log_mgmt

import "regexp"

/*
 * Function counts occurrences for the following HTTP methods: GET, PUT, DELETE, POST, HEAD.
 * Counts each unique IP address and URL(ignores queries)
 */
func CountLogMatchesIgnoresQuery(logChan <-chan string, urlCountChan chan<- map[string]int, ipCountChan chan<- map[string]int) {
	// match the IP address and URL(ignore query). Used https://regex101.com/.
	logRegex := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+).+(?:GET|POST|PUT|DELETE|HEAD)\s([^ ?]+)`)

	urlCount := make(map[string]int)
	ipCount := make(map[string]int)

	for log := range logChan {
		// Find matches in the log entry
		matches := logRegex.FindStringSubmatch(log)

		// Ensure the match is valid and has two groups
		if len(matches) == 3 {
			ip := matches[1]
			url := matches[2]

			urlCount[url]++
			ipCount[ip]++
		}
	}

	urlCountChan <- urlCount
	ipCountChan <- ipCount

	close(urlCountChan)
	close(ipCountChan)
}
