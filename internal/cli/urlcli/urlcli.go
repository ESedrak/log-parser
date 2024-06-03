package urlcli

import "github.com/spf13/cobra"

func NewURLCmd(urlCounts map[string]int, requestedNumURL int) *cobra.Command {
	var urlCmd = &cobra.Command{
		Use:   "url",
		Short: "URL related commands",
		Long:  "URL related commands",
	}

	urlCmd.AddCommand(newFindTopRequestedURLs(urlCounts, requestedNumURL))

	return urlCmd
}
