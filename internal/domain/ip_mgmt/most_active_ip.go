package ip_mgmt

import (
	"cmp"
	"errors"
	"log/slog"
	"slices"
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

	//sort the slice by counts in descending order
	slices.SortStableFunc(ipCount, func(a, b IPCount) int {
		if n := cmp.Compare(b.Count, a.Count); n != 0 {
			return n
		}
		// If counts are equal, order by IP (in ascending order)
		return cmp.Compare(a.IP, b.IP)
	})

	if len(ipCount) < requestedNum {
		slog.Warn("requested number exceeds the number of unique IPs: returning the maximum possible amount")
		return ipCount, nil
	}

	return ipCount[:requestedNum], nil
}
