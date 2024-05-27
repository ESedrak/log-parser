package ipcli

import (
	"log-parser/internal/domain/ip_mgmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func newFindUniqueIP() *cobra.Command {
	ipFindActiveCmd := &cobra.Command{
		Use:   "unique",
		Short: "Find the number of unique IPs",
		Long:  "Find the number of unique IPs within a log file",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			findUniquIP()
		},
	}

	return ipFindActiveCmd
}

func findUniquIP() {

	result, err := ip_mgmt.UniqueIPs(ipCounts)
	cobra.CheckErr(err)

	slog.Info("Unique IPs Count", "IPs", result)
}
