package logcli

import (
	"log-parser/internal/cli/ipcli"
	"log-parser/internal/cli/urlcli"
	"log-parser/internal/domain/log_mgmt"

	"github.com/spf13/cobra"
)

var ipCounts map[string]int
var urlCounts map[string]int

func NewLogCmd(logPath string, regex string) *cobra.Command {

	var logCmd = &cobra.Command{
		Use:   "log -- ip unique - ip active $(IP_COUNT) - url top $(URL_COUNT)",
		Short: "run multiple IP/URL commands separated by -",
		Long:  "Load Logs, Counts Log Matches, and can run multiple IP/URL commands separated by -",
		// Call the logHandlerFn to LoadLogs/CountLogMatches
		Run: func(cobraCmd *cobra.Command, args []string) {
			err := logHandlerFn(logPath, regex)
			cobra.CheckErr(err)
		},
		// After logs loaded/counted - add ability to invoke one or multiple subcommands for IP/URLs i.e: ./app/logparser log -- command1 - command2 - command3
		PostRun: func(cobraCmd *cobra.Command, args []string) {
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

	return logCmd
}

func logHandlerFn(logPath string, regex string) error {

	logChan := make(chan string)
	errChan := make(chan error)
	urlCountChan := make(chan map[string]int)
	ipCountChan := make(chan map[string]int)

	go log_mgmt.LoadLogs(logPath, logChan, errChan)

	go log_mgmt.CountLogMatch(regex, logChan, urlCountChan, ipCountChan)

	err := <-errChan
	if err != nil {
		return err
	}

	// receive URL counts
	urlCounts = <-urlCountChan

	// receive IP counts
	ipCounts = <-ipCountChan

	return nil
}
