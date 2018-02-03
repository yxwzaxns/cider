package server

import (
	"cider/config"
)

// Init xx
func Init() {
	// init CI/CD queue
	// CDQUEUE := make(chan string)
	c := new(config.Config)
	c.Init()
	c.InitEnvironmentVariable()
	r := NewRouter()
	r.Run(c.ListenIP + `:` + c.ListenPort)
}
