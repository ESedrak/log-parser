package ipcli

import "github.com/spf13/cobra"

func NewIP(ipCount map[string]int) *cobra.Command {
	var ipCmd = &cobra.Command{
		Use:   "IP",
		Short: "IP related commands",
		Long:  "IP related commands",
	}

	ipCmd.AddCommand(newFindMostActiveIP(ipCount), newFindUniqueIP(ipCount))

	return ipCmd
}
