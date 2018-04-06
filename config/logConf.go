package config

import (
	"os"

	"github.com/op/go-logging"
	"github.com/yxwzaxns/cider/utils"
)

var Log *logging.Logger

func LogInit() {
	Log = logging.MustGetLogger("cider")

	realLogFormat := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	fileLogFormat := logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05} %{longfile} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
	)
	realLogBackend := logging.NewLogBackend(os.Stderr, "", 0)
	realLogFormatter := logging.NewBackendFormatter(realLogBackend, realLogFormat)

	if utils.IsProduction() {
		logFile := utils.OpenFile(utils.GetEnv("CIDER_LOGPATH"))
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
