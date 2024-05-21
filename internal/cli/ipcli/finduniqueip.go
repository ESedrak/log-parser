package ipcli

import (
	"log-parser/internal/domain/ip_mgmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func newFindUniqueIP(ipCount map[string]int) *cobra.Command {
	ipFindActiveCmd := &cobra.Command{
		Use:   "find the number of unique IPs",
		Short: "Find the number of unique IPs",
		Long:  "Find the number of unique IPs within a log file",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			findUniquIP(ipCount)
		},
	}

	return ipFindActiveCmd
}

func findUniquIP(ipCount map[string]int) {

	result, err := ip_mgmt.UniqueIPs(ipCount)
	cobra.CheckErr(err)

	slog.Info("Unique IPs Count", "IPs", result)
}
