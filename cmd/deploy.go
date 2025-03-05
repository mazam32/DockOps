package cmd

import (
	"context"
	"fmt"
	"log"

	i "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
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

			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				log.Fatalf("Error creating Docker client: %v", err)
			}
			
			fmt.Println("Pulling Docker image:", image)
			_, err = cli.ImagePull(context.Background(), image, i.PullOptions{})

			if err != nil {
				log.Fatalf("Error pulling Docker image: %v", err)
			}
			fmt.Println("Image pulled successfully!")

		},
	}
}
