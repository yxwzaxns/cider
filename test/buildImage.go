package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver"
)

var DockerAPI float64 = 1.36

func createTar1() {
	if _, err := os.Stat("output.tar"); err == nil {
		os.RemoveAll("output.tar")
	}
	// path := "/tmp/cider_workspace/github_com_yxwzaxns_cider-ci-test"
	archiver.Tar.Make("output.tar", []string{"/tmp/cider_workspace/github_com_yxwzaxns_cider-ci-test/Dockerfile"})
}
func buildImage() int {

	os.Setenv("DOCKER_API_VERSION", strconv.FormatFloat(DockerAPI, 'f', 2, 64))
	CodeUrl := "github.com/yxwzaxns/cider-ci-test"
	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()
	if dockerBuildContext, err := os.Open("output.tar"); err == nil {
		defer dockerBuildContext.Close()
		cli, _ := client.NewEnvClient()
		options := types.ImageBuildOptions{
			Tags:           []string{CodeUrl + ":test3"},
			NoCache:        false,
			SuppressOutput: false,
			Remove:         true,
			ForceRemove:    true,
			PullParent:     true,
		}
		buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, options)
		if err != nil {
			fmt.Printf("%s", err.Error())
			DockerAPI -= 0.01
			os.Setenv("DOCKER_API_VERSION", strconv.FormatFloat(DockerAPI, 'f', 2, 64))
			return 1
		} else {
			// fmt.Printf("********* %s **********", buildResponse.OSType)
			response, err := ioutil.ReadAll(buildResponse.Body)
			if err != nil {

				fmt.Printf("%s", err.Error())
				return 1
			} else {
				// timeout := 5 * 60 * 1000
				// timeout := 30 * 1000
				// time.Sleep(time.Duration(timeout) * time.Millisecond)
				// cancel()
				fmt.Println(string(response))
			}
		}
	}
	return 0
}
func main() {
	createTar1()
	for {
		if buildImage() == 0 {
			break
		}
	}
}
