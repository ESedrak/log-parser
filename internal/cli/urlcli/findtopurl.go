package urlcli

import (
	"log-parser/internal/domain/url_mgmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

func newFindTopRequestedURLs(urlCount map[string]int) *cobra.Command {
	urlFindTopRequestedCmd := &cobra.Command{
		Use:   "find the top requested URLs",
		Short: "Find the top requested URLs",
		Long:  "Find the top requested URLs. \nRequested number of results is a required argument.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			intArg, err := strconv.Atoi(args[0])
			if err != nil {
				slog.Error("Error: argument must be an integer", "err", err)
				return
			}
			findTopRequestedURLs(urlCount, intArg)
		},
	}

	return urlFindTopRequestedCmd
}

func findTopRequestedURLs(urlCount map[string]int, requestedNum int) {

	result, err := url_mgmt.TopRequestedURLs(urlCount, requestedNum)
	cobra.CheckErr(err)

	slog.Info("Top Requested URLs:", "URLs", result)
}
