package ip_mgmt

import (
	"errors"
	"log/slog"
	"sort"
)

type IPCount struct {
	IP    string
	Count int
}

/*
 * Sorted by activity count in descending order
 * Returns the requested number of most active IP addresses (along with their counts)
 */
func MostActiveIP(ipCounts map[string]int, requestedNum int) ([]IPCount, error) {
	if requestedNum < 1 {
		return nil, errors.New("requested number for IPs has to be greater than 0")
	}

	var ipCount []IPCount

	// convert map entries into IPCount instances
	for ip, count := range ipCounts {
		ipCount = append(ipCount, IPCount{IP: ip, Count: count})
	}

	// sort the IPCount pairs based on count
	sort.Slice(ipCount, func(i, j int) bool {
		return ipCount[j].Count < ipCount[i].Count
	})

	if len(ipCount) < requestedNum {
		slog.Warn("requested number is less than the number of unique IPs: returning the maximum possible amount")
		return ipCount, nil
	}

	return ipCount[:requestedNum], nil
}
