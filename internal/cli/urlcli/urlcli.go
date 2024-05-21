package urlcli

import "github.com/spf13/cobra"

func NewURL(urlCount map[string]int) *cobra.Command {
	var urlCmd = &cobra.Command{
		Use:   "url",
		Short: "URL related commands",
		Long:  "URL related commands",
	}

	urlCmd.AddCommand(newFindTopRequestedURLs(urlCount))

	return urlCmd
}
