package log_mgmt

import (
	"log-parser/internal/ip_mgmt"
	"log-parser/internal/url_mgmt"
	"strings"
)

/*
 * Question: Will you always want to remove the query?
 * Question: Will the wanted HTTP methods ever change?
 *
 * Function counts occurrences for the following HTTP methods: GET, PUT, DELETE, POST, HEAD.
 * Counts each unique IP address and URLs(ignores queries) in log matches
 */
func CountLogMatchesIgnoresQuery(logMatches [][]string) (*url_mgmt.URLCounter, *ip_mgmt.IPCounter) {
	urlCounts := url_mgmt.NewUrlCounts()
	ipCounts := ip_mgmt.NewIPCounter()

	for _, match := range logMatches {
		ip := match[0]         // IP address is stored at index 0
		httpMethod := match[2] // HTTPMethod is stored at index 2
		url := match[3]        // URL is stored at index 3

		if !isValidHTTPMethod(httpMethod) {
			continue // skip the current iteration as not a wanted HTTP method
		}

		url = strings.Split(url, "?")[0] // remove query from URL before counting it

		urlCounts.URLCounts[url]++
		ipCounts.IPCounts[ip]++
	}

	return urlCounts, ipCounts
}

func isValidHTTPMethod(httpMethod string) bool {
	validHTTPMethods := map[string]bool{
		"GET":    true,
		"PUT":    true,
		"DELETE": true,
		"POST":   true,
		"HEAD":   true,
	}
	_, ok := validHTTPMethods[httpMethod]
	return ok
}
