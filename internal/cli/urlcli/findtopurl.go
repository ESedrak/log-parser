package urlcli

import (
	"log-parser/internal/domain/url_mgmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

func newFindTopRequestedURLs(urlCounts map[string]int, requestedNumURL int) *cobra.Command {
	urlFindTopRequestedCmd := &cobra.Command{
		Use:   "top",
		Short: "Return the top requested URLs",
		Long:  "Return the top requested URLs. \nSpecify the requested number of results to return. Otherwise it will default to the configs requestedNum for URLs",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var requestedNum int
			var err error

			if len(args) == 0 {
				requestedNum = requestedNumURL
			} else {
				requestedNum, err = strconv.Atoi(args[0])
				if err != nil {
					slog.Error("Requested number of URLs to return must be an integer", "err", err)
					return
				}
			}
			findTopRequestedURLs(urlCounts, requestedNum)
		},
	}

	return urlFindTopRequestedCmd
}

func findTopRequestedURLs(urlCounts map[string]int, requestedNum int) {

	result, err := url_mgmt.TopRequestedURLs(urlCounts, requestedNum)
	cobra.CheckErr(err)

	slog.Info("Top Requested URLs:", "URLs", result)
}
