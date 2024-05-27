package ipcli

import "github.com/spf13/cobra"

var ipCounts map[string]int

func NewIPCmd(ipCount map[string]int) *cobra.Command {
	ipCounts = ipCount
	var ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "IP related commands",
		Long:  "IP related commands",
	}

	ipCmd.AddCommand(newFindMostActiveIP(), newFindUniqueIP())

	return ipCmd
}
