package global

import (
	"cider/core"
	"cider/db"
	"path/filepath"
)

var (
	BasePath   string
	Projects   db.ProjectTable
	EventsChan chan string
	Core       *core.Core
)

func Init() {
	path, _ := filepath.Abs(".")
	BasePath = path
	EventsChan = make(chan string, 10)
	Core = new(core.Core)
	Core.Init(EventsChan)
}
