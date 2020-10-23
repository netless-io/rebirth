package run

import (
	"github.com/netless-io/rebirth/internal/app/rebirth/chrome"
	"github.com/netless-io/rebirth/internal/app/rebirth/server"
	"github.com/netless-io/rebirth/internal/pkg/cobrahelp"
	"github.com/netless-io/rebirth/internal/pkg/logs"
	"github.com/spf13/cobra"
)

var log = logs.Rebirth

// NewRunCommand creates a new `rebirth run` command
func NewRunCommand() *cobra.Command {
	var flags runFlags
	flags.Chrome.Init()

	cmd := &cobra.Command{
		Use:                   "run [OPTIONS] URL",
		Short:                 "Run a recording task",
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
		TraverseChildren:      true,
		Args:                  cobrahelp.RequiresMinArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			flags.Record.UrlAddress = args[0]

			printCommandParams(&flags)

			cmd, err := chrome.Cmd(&flags.Chrome)
			if err != nil {
				return err
			}

			if err := cmd.Start(); err != nil {
				return err
			}

			server.RecordInfoAPI(&flags.Record, cmd.Process)
			go server.Listen(cmd.Process)

			return cmd.Wait()
		},
	}

	cmd.Flags().StringVarP(&flags.Record.StartTime, "start-time", "S", "", "When to start recording")
	cmd.Flags().StringVarP(&flags.Record.EndTime, "end-time", "E", "", "When to stop recording")
	cmd.Flags().StringVarP(&flags.Record.Filename, "output-filename", "o", "rebirth", "After recording, the output video file name")
	cmd.Flags().StringVarP(&flags.RuntimeEnv, "runtime-platform", "p", "docker", "Run the recording task in that runtime platform {docker | local}")
	cmd.Flags().StringVarP(&flags.Chrome.Path, "chrome-path", "", flags.Chrome.Path, "Chrome executable file path")
	cmd.Flags().StringVarP(&flags.Chrome.UserDataDir, "chrome-data-dir", "", flags.Chrome.UserDataDir, "Chrome user data dir")
	cmd.Flags().StringVarP(&flags.Chrome.ExtensionDir, "extension-dir", "", flags.Chrome.ExtensionDir, "Chrome extension dir")
	cmd.Flags().IntVarP(&flags.Record.FPS, "fps", "F", 30, "FPS during recording")

	cmd.SetFlagErrorFunc(cobrahelp.FlagErrorFunc)

	return cmd
}

// printCommandParams output run flags
func printCommandParams(f *runFlags) {
	log.Debugf("url address is: %s", f.Record.UrlAddress)
	log.Debugf("start record time is: %s", f.Record.StartTime)
	log.Debugf("end record time is: %s", f.Record.EndTime)
	log.Debugf("output video filename is: %s", f.Record.Filename)
	log.Debugf("current runtime environment is: %s", f.RuntimeEnv)
	log.Debugf("chrome executable file path is: %s", f.Chrome.Path)
	log.Debugf("chrome user data dir is: %s", f.Chrome.UserDataDir)
	log.Debugf("chrome extension dir is: %s", f.Chrome.ExtensionDir)
}
