package ipcli

import (
	"log-parser/internal/domain/ip_mgmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

func newFindMostActiveIP(ipCount map[string]int) *cobra.Command {
	ipFindActiveCmd := &cobra.Command{
		Use:   "active",
		Short: "Find the most active IP",
		Long:  "Find the most active IP. \nRequested number of results is a required argument.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			intArg, err := strconv.Atoi(args[0])
			if err != nil {
				slog.Error("Error: argument must be an integer", "err", err)
				return
			}
			findMostActiveIP(ipCount, intArg)
		},
	}

	return ipFindActiveCmd
}

func findMostActiveIP(ipCount map[string]int, requestedNum int) {

	result, err := ip_mgmt.MostActiveIP(ipCount, requestedNum)
	cobra.CheckErr(err)

	slog.Info("Most Active IPs: ", "IPs", result)
}
