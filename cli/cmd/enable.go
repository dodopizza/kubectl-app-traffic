package cmd

import (
	"fmt"

	"github.com/dodopizza/kubectl-app-traffic/cli/internal"
	"github.com/spf13/cobra"
)

func Enable() *cobra.Command {
	options := &internal.ToggleOptions{}
	command := &cobra.Command{
		Use:   "enable [service]",
		Short: "Enable traffic back to k8s service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Service = args[0]
			patch := &internal.Patch{
				Operation: "remove",
				Path:      "/spec/selector/offline",
			}

			fmt.Printf("Enable traffic to %s/%s\n", *options.Kube.Namespace, options.Service)
			err := options.Patch(patch)
			if err != nil {
				fmt.Printf("Failed to enable traffic to %s/%s\nTraffic may be already enabled for this service\n",
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
