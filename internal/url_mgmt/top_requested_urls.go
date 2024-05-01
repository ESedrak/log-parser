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
 * Question: How to handle excess occurrences?
 * E.g. if requestCount is 3, but there are 5 logs with the same number of URL calls. Only 3 results are returned.
 * Method is not idempotent
 *
 * Returns the top N most requested URL's along with their counts
 * Sorted by activity count in descending order
 */
func (u *URLCounter) TopRequestedURLs(requestedCount int) ([]URLCount, error) {
	if requestedCount < 1 {
		return nil, errors.New("requested count for URLs has to be greater than 0")
	}

	var urlCount []URLCount

	// convert map entries into URLCount instances
	for url, count := range u.URLCounts {
		urlCount = append(urlCount, URLCount{URL: url, Count: count})
	}

	// sort the slice by counts in descending order
	sort.Slice(urlCount, func(i, j int) bool {
		return urlCount[j].Count < urlCount[i].Count
	})

	if len(urlCount) < requestedCount {
		//do we want this warning here?
		slog.Warn("requested number of URLs is less than the number of unique URLs")
		return urlCount, nil
	}

	return urlCount[:requestedCount], nil
}
