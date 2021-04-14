package main

import (
	"os"

	"github.com/dodopizza/kubectl-app-traffic/cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	command := &cobra.Command{
		Use:          "kubectl-app_traffic",
		Short:        "Disable or enable traffic to k8s service",
		SilenceUsage: true,
	}
	command.AddCommand(cmd.Enable())
	command.AddCommand(cmd.Disable())
	command.AddCommand(cmd.Version())

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
