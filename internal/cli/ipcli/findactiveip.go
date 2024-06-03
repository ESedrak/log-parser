package ipcli

import (
	"log-parser/internal/domain/ip_mgmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

func newFindMostActiveIP(ipCounts map[string]int, requestedNumIP int) *cobra.Command {
	ipFindActiveCmd := &cobra.Command{
		Use:   "active",
		Short: "Return the most active IPs",
		Long:  "Return the most active IPs. \nSpecify the requested number of results to return. Otherwise it will default to the configs requestedNum for IPs",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var requestedNum int
			var err error

			if len(args) == 0 {
				requestedNum = requestedNumIP
			} else {
				requestedNum, err = strconv.Atoi(args[0])
				if err != nil {
					slog.Error("Requested number of IPs to return must be an integer", "err", err)
					return
				}
			}
			findMostActiveIP(ipCounts, requestedNum)
		},
	}

	return ipFindActiveCmd
}

func findMostActiveIP(ipCounts map[string]int, requestedNum int) {

	result, err := ip_mgmt.MostActiveIP(ipCounts, requestedNum)
	cobra.CheckErr(err)

	slog.Info("Most Active IPs:", "IPs", result)
}
