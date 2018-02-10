package core

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	yaml "gopkg.in/yaml.v2"
)

func StartCD(url string, msgChan chan M) {
	// imageName := url + ":" + "latest"
	if WorkPath == "" {
		basePath := filepath.Join(RootDir, "cider_workspace")
		WorkPath = filepath.Join(basePath, convertURLToPath(url))
	}
	ctx := context.Background()

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	composeFilePath := filepath.Join(WorkPath, "docker-compose.yml")
	composeFile, err := ioutil.ReadFile(composeFilePath)
	if err != nil {
		panic(err)
	}

	c := DockerCompose{}
	err = yaml.Unmarshal(composeFile, &c)
	if err != nil {
		panic(err)
	}
	for i := range c.Services {
		containerName := IToC(c.Services[i].ImageName)
		// CheckContainerExist
		// stop and remove it if it exist
		if id := CheckContainerExist(containerName); id != nil {
			timeout := time.Duration(10 * time.Second)
			cli.ContainerStop(ctx, id.(string), &timeout)
			cli.ContainerRemove(ctx, id.(string), types.ContainerRemoveOptions{
				Force: true,
			})
		}
		// end CheckContainerExist
		portsMap := nat.PortMap{}
		for j := range c.Services[i].Ports {
			portInfo := ParsePort(c.Services[i].Ports[j])
			portsBinding := []nat.PortBinding{{HostIP: portInfo[0][0], HostPort: portInfo[0][1]}}
			portsMap[nat.Port(portInfo[1][0])] = portsBinding
		}
		hostConfig := &container.HostConfig{
			PortBindings: portsMap,
		}

		// container config set
		exposePortInfo := nat.PortSet{}
		for j := range c.Services[i].Ports {
			portInfo := ParsePort(c.Services[i].Ports[j])
			exposePortInfo[nat.Port(portInfo[1][0])] = struct{}{}
		}

		containerConfig := &container.Config{
			Image:        c.Services[i].ImageName,
			ExposedPorts: exposePortInfo,
		}
		resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, containerName)
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}
	}
}

func CheckContainerExist(name string) interface{} {
	name = "/" + name
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}
	for _, c := range containerList {
		if c.Names[0] == name {
			return c.ID
		}
	}
	return nil
}

func IToC(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(name))[0:17]
}
