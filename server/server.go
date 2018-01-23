package server

import (
	"cider/config"
)

// Init xx
func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.ListenIP + `:` + c.ListenPort)
}
