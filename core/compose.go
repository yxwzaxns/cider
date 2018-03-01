package core

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	ImageName   string `yaml:"image"`
	Ports       []string
	Restart     string
	Privileged  bool
	Volumes     []string
	Environment []string
}
type DockerCompose struct {
	Version  string
	Services map[string]Server
}

func ComposeParse(composeFile string) DockerCompose {
	t := DockerCompose{}
	file, _ := ioutil.ReadFile(composeFile)
	fmt.Println(file)
	err := yaml.Unmarshal(file, &t)
	if err != nil {
		panic(err)
	}
	return t
}
