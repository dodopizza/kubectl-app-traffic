package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	VersionTag = "undefined"
)

func Version() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print CLI version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(VersionTag)
		},
	}
	return command
}
