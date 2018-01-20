package server

import (
	"cider/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.ListenIp + `:` + c.ListenPort)
}
