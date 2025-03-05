package main

import (
	"fmt"
	"os"

	"github.com/mazam32/DockOps/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "DockOps",
		Short: "DockOps is a CLI for automated deployments",
		Long:  "DockOps helps in deploying applications using Docker and SSH, making remote execution seamless.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to DockOps! Use --help to see available commands.")
		},
	}

	rootCmd.AddCommand(cmd.NewDeployCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
