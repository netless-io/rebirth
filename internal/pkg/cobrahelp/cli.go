package cobrahelp

import (
	"fmt"
	"github.com/netless-io/rebirth/internal/app/rebirth/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func FlagErrorFunc(cmd *cobra.Command, err error) error {
	if err == nil {
		return nil
	}

	usage := ""
	if cmd.HasSubCommands() {
		usage = "\n\n" + cmd.UsageString()
	}
	return cli.StatusError{
		Status:     fmt.Sprintf("%s\nSee '%s --help'%s", err, cmd.CommandPath(), usage),
		StatusCode: 125,
	}
}

// RequiresMinArgs returns an error if there is not at least min args
func RequiresMinArgs(min int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) >= min {
			return nil
		}
		return errors.Errorf(
			"%q requires at least %d %s.\nSee '%s --help'.\n\nUsage:  %s\n\n%s",
			cmd.CommandPath(),
			min,
			pluralize("argument", min),
			cmd.CommandPath(),
			cmd.UseLine(),
			cmd.Short,
		)
	}
}

//nolint: unparam
func pluralize(word string, number int) string {
	if number == 1 {
		return word
	}
	return word + "s"
}
