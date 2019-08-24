package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/heroku/docker-registry-client/registry"
	"os"
)

func getContainer(cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		for _, container := range containers {
			fmt.Println("Container ID:", container.ID, " Current image: ", container.Image)
		}
	} else {
		fmt.Println("There are no containers running")
	}
}

func initLocalDocker() *client.Client {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)
	return cli
}

func initRegistryDocker() *registry.Registry {
	urlRegistry := os.Args[1]
	username := ""
	password := ""

	if len(os.Args) >= 3 {
		username = os.Args[2]
	}
	if len(os.Args) >= 4 {
		password = os.Args[2]
	}
	hub, err := registry.New(urlRegistry, username, password)
	if err != nil {
		panic(err)
	}
	return hub
}
