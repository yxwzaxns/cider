package global

import (
	"cider/config"
	"cider/core"
	"cider/db"
	"cider/utils"
	"os"
	"path/filepath"

	"github.com/op/go-logging"
)

var (
	BasePath   string
	Projects   db.ProjectTable
	EventsChan chan string
	Core       *core.Core
	Config     *config.Config
	Log        *logging.Logger
)

func Init() {
	path, _ := filepath.Abs(".")
	BasePath = path
	EventsChan = make(chan string, 10)
	Core = new(core.Core)
	Core.Init(EventsChan)

	loggingInit()
}

func loggingInit() {
	Log = logging.MustGetLogger("cider")

	realLogFormat := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	fileLogFormat := logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05} %{longfile} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
	)
	realLogBackend := logging.NewLogBackend(os.Stderr, "", 0)
	realLogFormatter := logging.NewBackendFormatter(realLogBackend, realLogFormat)

	if Config.Debug == false {
		logFile := utils.OpenFile(Config.LogPath)
		fileLogBackend := logging.NewLogBackend(logFile, "", 0)
		fileLogFormatter := logging.NewBackendFormatter(fileLogBackend, fileLogFormat)
		logging.SetBackend(realLogFormatter, fileLogFormatter)
	} else {
		logging.SetBackend(realLogFormatter)
	}
	// log.Debugf("debug %s", "secret")
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")
}

func ImportConfig(c *config.Config) {
	Config = c
}
