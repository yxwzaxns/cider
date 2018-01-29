package server

import (
	"cider/config"
)

// Init xx
func Init() {
	// init CI/CD queue
	// CDQUEUE := make(chan string)
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.ListenIP + `:` + c.ListenPort)
}
