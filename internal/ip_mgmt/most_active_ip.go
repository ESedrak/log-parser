package ip_mgmt

import (
	"errors"
	"sort"

	"github.com/sirupsen/logrus"
)

type IPCount struct {
	IP    string
	Count int
}

/*
 * Question: How to handle excess occurrences?
 * E.g. if requestCount is 3, but there are 5 logs with the same number of IP calls. Only 3 results are returned.
 * Method is not idempotent
 *
 * Returns the top N most active IP addresses along with their counts
 * Sorted by activity count in descending order
 */
func (i *IPCounter) MostActiveIP(requestedCount int) ([]IPCount, error) {
	if requestedCount < 1 {
		return nil, errors.New("requested count for IPs has to be greater than 0")
	}

	var ipCount []IPCount

	// convert map entries into IPCount instances
	for ip, count := range i.IPCounts {
		ipCount = append(ipCount, IPCount{IP: ip, Count: count})
	}

	// sort the IPCount pairs based on count
	sort.Slice(ipCount, func(i, j int) bool {
		return ipCount[j].Count < ipCount[i].Count
	})

	if len(ipCount) < requestedCount {
		//do we want this warning here?
		logrus.Warn("requested number of IPs is less than the number of unique IPs")
		return ipCount, nil
	}

	return ipCount[:requestedCount], nil
}
