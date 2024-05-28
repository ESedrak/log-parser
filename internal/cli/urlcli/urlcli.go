package urlcli

import "github.com/spf13/cobra"

var urlCounts map[string]int

func NewURLCmd(urlCount map[string]int) *cobra.Command {
	urlCounts = urlCount
	var urlCmd = &cobra.Command{
		Use:   "url",
		Short: "URL related commands",
		Long:  "URL related commands",
	}

	urlCmd.AddCommand(newFindTopRequestedURLs())

	return urlCmd
}
