package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func CheckContainerExist(name string) interface{} {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	for _, c := range containerList {
		if c.Image == name {
			return c.ID
		}
	}
	return nil
}

func main() {
	imageName := "alpine-dev"
	os.Setenv("DOCKER_API_VERSION", "1.35")
	if id := CheckContainerExist(imageName); id != nil {
		fmt.Printf("%s", id)
	}
}
