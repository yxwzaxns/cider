package main

import (
	"context"
	"io/ioutil"

	"github.com/docker/docker/client"
)

func main() {
	cli, _ := client.NewEnvClient()
	// println(client.DefaultDockerHost)
	// u, _ := cli.DiskUsage(context.Background())
	// 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	dockerfile, _ := ioutil.ReadFile("Dockerfile")
	println()
	// 	for _, container := range containers {
	// 		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	// 	}
	// dockerfile, _ := ioutil.ReadFile("Dockerfile")
	// fmt.Printf("File contents: %s", dockerfile)
	cli.ImageBuild(context.Background(), dockerfile)
}
