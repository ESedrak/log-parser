package ipcli

import (
	"log-parser/internal/domain/ip_mgmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func newFindUniqueIP(ipCounts map[string]int) *cobra.Command {
	ipFindActiveCmd := &cobra.Command{
		Use:   "unique",
		Short: "Return the number of unique IPs",
		Long:  "Return the number of unique IPs within a log file",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			findUniquIP(ipCounts)
		},
	}

	return ipFindActiveCmd
}

func findUniquIP(ipCounts map[string]int) {

	result, err := ip_mgmt.UniqueIPs(ipCounts)
	cobra.CheckErr(err)

	slog.Info("Unique IPs Count", "IPs", result)
}
