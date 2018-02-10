package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	ImageName string `yaml:"image"`
	Ports     []string
}
type DockerCompose struct {
	Version  string
	Services map[string]Server
}

func ComposeParse(composeFile string) DockerCompose {
	t := DockerCompose{}
	file, _ := ioutil.ReadFile(composeFile)
	err := yaml.Unmarshal(file, &t)
	if err != nil {
		panic(err)
	}
	return t
}
