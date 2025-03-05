package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeployCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploys an application using Docker",
		Args:  cobra.ExactArgs(1),
		Long:  "Deploys a given application by pulling its Docker image and running it on the server.",
		Run: func(cmd *cobra.Command, args []string) {
			image := args[0]
			fmt.Println("Deploying:", image)

		},
	}
}
