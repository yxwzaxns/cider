package main

import (
	"cider/config"
	"cider/db"
	G "cider/global"
	"cider/server"
)

func main() {
	// ProjectPath, _ = filepath.Abs("./")
	c := new(config.Config)
	c.Init()
	G.Init()
	G.ImportConfig(c)
	db.Init(G.Config.AppDbPath)

	server.Init()
}
