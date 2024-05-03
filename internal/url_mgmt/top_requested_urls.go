package url_mgmt

import (
	"errors"
	"log/slog"
	"sort"
)

type URLCount struct {
	URL   string
	Count int
}

/*
 * Returns the top N most requested URL's along with their counts
 * Sorted by activity count in descending order
 */
func TopRequestedURLs(urlCounts map[string]int, requestedNum int) ([]URLCount, error) {
	if requestedNum < 1 {
		return nil, errors.New("requested count for URLs has to be greater than 0")
	}

	var urlCount []URLCount

	// convert map entries into URLCount instances
	for url, count := range urlCounts {
		urlCount = append(urlCount, URLCount{URL: url, Count: count})
	}

	// sort the slice by counts in descending order
	sort.Slice(urlCount, func(i, j int) bool {
		return urlCount[j].Count < urlCount[i].Count
	})

	if len(urlCount) < requestedNum {
		slog.Warn("requested number is less than the number of unique URLs: returning the maximum possible amount")
		return urlCount, nil
	}

	return urlCount[:requestedNum], nil
}
