package rebirth

import (
	"fmt"
	"github.com/netless-io/rebirth/internal/app/rebirth/cli"
	"github.com/netless-io/rebirth/internal/app/rebirth/cli/command/run"
	"github.com/netless-io/rebirth/internal/pkg/logs"
	"github.com/spf13/cobra"
	"os"
)

var log = logs.Server

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "rebirth COMMAND [OPTIONS]",
		Short:                 "Record website changes in the background",
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
		TraverseChildren:      true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			setLogParams()
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				_ = cmd.Help()
				return
			}
		},
	}

	cmd.AddCommand(run.NewRunCommand())

	cmd.PersistentFlags().BoolVarP(&cli.RootFlags.Debug, "debug", "D", false, "Enable debug mode")
	cmd.PersistentFlags().StringVarP(&cli.RootFlags.LogFormat, "log-format", "", "json", "Set log output format {json | text}")

	return cmd
}

func setLogParams() {
	if cli.RootFlags.Debug {
		logs.EnableDebug()
	}

	logs.SetFormatter(cli.RootFlags.LogFormat)
}

func Execute() {
	var command = newRootCommand()
	if err := command.Execute(); err != nil {
		if stderr, ok := err.(cli.StatusError); ok {
			if stderr.Status != "" {
				log.Error(stderr.Status)
			}

			// StatusError should only be used for errors, and all errors should
			// have a non-zero exit status, so never exit with 0
			if stderr.StatusCode == 0 {
				os.Exit(1)
			}

			os.Exit(stderr.StatusCode)
		}

		fmt.Println(err.Error())
		os.Exit(1)
	}
}
