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
		Use:   "disable service|ingress [service_name|ingress_name]",
		Short: "Disable traffic from k8s service or ingress",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var patch *internal.Patch
			switch args[0] {
			case "service":
				options.Service = args[1]
				fmt.Println(args)
				patch = &internal.Patch{
					Operation: "add",
					Path:      "/spec/selector/offline",
					Value:     strings.Replace(time.Now().UTC().Format(time.RFC3339), ":", ".", -1),
				}

				fmt.Printf("Disable traffic from %s/%s\n", *options.Kube.Namespace, options.Service)
			case "ingress":
				options.Ingress = args[1]
				patch = &internal.Patch{
					Operation: "add",
					Path:      "/metadata/annotations/nginx.ingress.kubernetes.io~1whitelist-source-range",
					Value:     "127.0.0.1/32",
				}
				fmt.Printf("Disable traffic from %s/%s\n", *options.Kube.Namespace, options.Ingress)
			default:
				fmt.Printf("Unknown resource")
				return nil
			}
			err := options.Patch(patch)
			if err != nil {
				fmt.Printf("Failed to disable traffic from %s/%s\n",
					*options.Kube.Namespace,
					args[1])
			}
			return err
		},
	}

	command.
		Flags().
		AddFlagSet(options.Parse())

	return command
}
