package cli

import (
	"log-parser/internal/cli/logcli"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {

	var logparser = &cobra.Command{
		Use:   "logparser",
		Short: "LogParser CLI application",
		Long:  "Logparser CLI application\nUse this CLI for all your log parsing needs.",
	}
	logparser.AddCommand(logcli.NewLogCmd())

	return logparser
}
