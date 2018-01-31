package main

import (
	G "cider/global"
	"cider/server"
)

func main() {
	// ProjectPath, _ = filepath.Abs("./")
	G.Init()
	server.Init()
}
