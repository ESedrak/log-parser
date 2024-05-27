package logcli

import (
	"log-parser/internal/cli/ipcli"
	"log-parser/internal/cli/urlcli"
	"log-parser/internal/config"
	"log-parser/internal/domain/log_mgmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func init() {
	config.Init("./config", "config", "json")
}

var ipCounts map[string]int
var urlCounts map[string]int

func NewLogCmd() *cobra.Command {

	// Ability to invoke one or multiple commands i.e. logparser and -- command1 - command2 - command3
	var andCmd = &cobra.Command{
		Use:   "and -- ip active $(IP_COUNT) - ip unique - url top $(URL_COUNT)",
		Short: "run multiple logparser commands separated by -",
		Long:  "Load Logs, Counts Log Matches, and allows to run multiple logparser commands separated by -",
		Run: func(cobraCmd *cobra.Command, args []string) {
			logHandlerFn()

			var cmdParts []string
			var cmdList [][]string
			for _, arg := range args {
				if arg == "-" {
					if len(cmdParts) > 0 {
						cmdList = append(cmdList, cmdParts)
						cmdParts = []string{}
					}
				} else {
					cmdParts = append(cmdParts, arg)
				}
			}
			cmdList = append(cmdList, cmdParts)

			for _, cmdParts := range cmdList {
				cmd := &cobra.Command{}
				//Add all IP/URL commands
				cmd.AddCommand(ipcli.NewIPCmd(ipCounts), urlcli.NewURLCmd(urlCounts))
				cmd.SetArgs(cmdParts)
				cmd.Execute()
			}
		},
	}

	return andCmd
}

func logHandlerFn() {
	cfg := config.Values

	logChan := make(chan string)
	errChan := make(chan error)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(cfg.Path.LogPath, logChan, errChan)

	go log_mgmt.CountLogMatch(cfg.Regex.MatchIPsURLsIgnoreQuery, logChan, urlCountChan, ipCountChan)

	err := <-errChan
	if err != nil {
		slog.Error("error", "err", err)
		return
	}

	// receive URL counts
	urlCounts = <-urlCountChan

	// receive IP counts
	ipCounts = <-ipCountChan
}
