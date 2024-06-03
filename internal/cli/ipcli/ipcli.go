package ipcli

import "github.com/spf13/cobra"

func NewIPCmd(ipCounts map[string]int, requestedNumIP int) *cobra.Command {
	var ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "IP related commands",
		Long:  "IP related commands",
	}

	ipCmd.AddCommand(newFindMostActiveIP(ipCounts, requestedNumIP), newFindUniqueIP(ipCounts))

	return ipCmd
}
