package global

import (
	"cider/core"
	"cider/db"
)

var (
	BasePath   string
	Projects   db.ProjectTable
	EventsChan chan string
	Core       *core.Core
)

func Init() {
	BasePath = ""
	EventsChan = make(chan string, 10)
	Core = new(core.Core)
	Core.Init(EventsChan)
}
