package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/dodopizza/kubectl-app-traffic/cli/internal"
	"github.com/spf13/cobra"
)

func Disable() *cobra.Command {
	options := &internal.ToggleOptions{}
	command := &cobra.Command{
		Use:   "disable [service]",
		Short: "Disable traffic from k8s service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Service = args[0]
			patch := &internal.Patch{
				Operation: "add",
				Path:      "/spec/selector/offline",
				Value:     strings.Replace(time.Now().UTC().Format(time.RFC3339), ":", ".", -1),
			}

			fmt.Printf("Disable traffic from %s/%s\n", *options.Kube.Namespace, options.Service)
			err := options.Patch(patch)
			if err != nil {
				fmt.Printf("Failed to disable traffic from %s/%s\n",
					*options.Kube.Namespace,
					options.Service)
			}
			return err
		},
	}

	command.
		Flags().
		AddFlagSet(options.Parse())

	return command
}
