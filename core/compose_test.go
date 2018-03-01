package core

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// version: '3'
// services:
// 	myblog:
// 		image: daocloud.io/yxwzaxns/aong-ghost:master-fc6fdb7
// 		privileged: false
// 		restart: always
// 		ports:
// 		- 127.0.0.1:2368:2368
// 		volumes:
// 		- /vo:/ghost/content
// 		environment:
// 		- NODE_ENV=production
// 		- PROTO_TYPE=https

func TestParseDockerCompose(t *testing.T) {
	path := "/tmp/cider_workspace/github_com_yxwzaxns_cider-ci-test/docker-compose.yml"

	c := ComposeParse(path)

	for _, app := range c.Services {
		spew.Dump(app)
	}
}
