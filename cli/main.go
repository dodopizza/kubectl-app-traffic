package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	Version = "undefined"
)

func enable() *cobra.Command {
	options := &ToggleOptions{}
	command := &cobra.Command{
		Use:   "enable [service]",
		Short: "Enable traffic back to k8s service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Service = args[0]
			patch := &Patch{
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

func disable() *cobra.Command {
	options := &ToggleOptions{}
	command := &cobra.Command{
		Use:   "disable [service]",
		Short: "Disable traffic from k8s service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Service = args[0]
			patch := &Patch{
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

func version() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print CLI version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
	return command
}

func main() {
	command := &cobra.Command{
		Use:          "kubectl-app-traffic",
		Short:        "Disable or enable traffic to k8s service",
		SilenceUsage: true,
	}

	command.AddCommand(enable())
	command.AddCommand(disable())
	command.AddCommand(version())

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
