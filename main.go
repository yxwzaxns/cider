package main

import (
	"cider/config"
	"cider/db"
	G "cider/global"
	"cider/server"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// ProjectPath, _ = filepath.Abs("./")

	app := cli.NewApp()
	app.Name = "cider"
	app.Usage = "Cider is a concise CI/CD tool for docker!"
	app.Action = func(c *cli.Context) error {
		conf := new(config.Config)
		conf.Init()
		G.Init()
		G.ImportConfig(conf)
		db.Init(G.Config.AppDbPath)

		server.Init()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
