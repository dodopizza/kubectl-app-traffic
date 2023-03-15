package cmd

import (
	"fmt"

	"github.com/dodopizza/kubectl-app-traffic/cli/internal"
	"github.com/spf13/cobra"
)

func Enable() *cobra.Command {
	options := &internal.ToggleOptions{}
	command := &cobra.Command{
		Use:   "enable service|ingress [service_name|ingress_name]",
		Short: "Enable traffic back to k8s service or ingress",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var patch *internal.Patch
			switch args[0] {
			case "service":
				options.Service = args[1]
				patch = &internal.Patch{
					Operation: "remove",
					Path:      "/spec/selector/offline",
				}

				fmt.Printf("Enable traffic to %s/%s\n", *options.Kube.Namespace, options.Service)
			case "ingress":
				options.Ingress = args[1]
				patch = &internal.Patch{
					Operation: "remove",
					Path:      "/metadata/annotations/nginx.ingress.kubernetes.io~1whitelist-source-range",
				}
				fmt.Printf("Enable traffic to %s/%s\n", *options.Kube.Namespace, options.Ingress)
			default:
				fmt.Printf("Unknown resource")
				return nil
			}

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
