package core

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func TestStartCD(t *testing.T) {
	mChan := make(chan M, 10)
	imageName := "github.com/yxwzaxns/cider-ci-test"
	os.Setenv("DOCKER_API_VERSION", "1.35")
	println("start test cd")
	StartCD(imageName, mChan)
}

func TestStopAndRemoveContainer(t *testing.T) {
	os.Setenv("DOCKER_API_VERSION", "1.35")
	imageName := "github.com/yxwzaxns/cider-ci-test:latest"
	containerName := IToC(imageName)
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	// CheckContainerExist
	// stop and remove it if it exist
	if id := CheckContainerExist(containerName); id != nil {
		timeout := time.Duration(10 * time.Second)
		cli.ContainerStop(ctx, id.(string), &timeout)
		cli.ContainerRemove(ctx, id.(string), types.ContainerRemoveOptions{
			Force: true,
		})
	} else {
		println("no container be find")
	}
}
