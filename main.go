package main

import (
	"log"
	"os"

	"github.com/yxwzaxns/cider/db"
	"github.com/yxwzaxns/cider/global"
	"github.com/yxwzaxns/cider/server"

	"github.com/yxwzaxns/cider/config"

	"github.com/urfave/cli"
)

func main() {
	// ProjectPath, _ = filepath.Abs("./")

	app := cli.NewApp()
	app.Name = "cider"
	app.Usage = "Cider is a concise CI/CD tool for docker!"
	app.Action = func(c *cli.Context) error {
		global.Init()
		config.Init()
		db.Init(config.Conf.AppDbPath)
		server.Init()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
