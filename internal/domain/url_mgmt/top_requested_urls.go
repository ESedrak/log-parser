package url_mgmt

import (
	"cmp"
	"errors"
	"log/slog"
	"slices"
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

	//sort the slice by counts in descending order
	slices.SortStableFunc(urlCount, func(a, b URLCount) int {
		if n := cmp.Compare(b.Count, a.Count); n != 0 {
			return n
		}
		// If counts are equal, order by URL (in ascending order)
		return cmp.Compare(a.URL, b.URL)
	})

	if len(urlCount) < requestedNum {
		slog.Warn("requested number exceeds the the number of unique URLs: returning the maximum possible amount")
		return urlCount, nil
	}

	return urlCount[:requestedNum], nil
}
