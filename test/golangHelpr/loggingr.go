package main

import (
	"io"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var realLogFormat = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var fileLogFormat = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{longfile} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
)

func openFile(path string) io.Writer {
	if f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	} else {
		return f
	}
	return nil
}
func main() {
	// For demo purposes, create two backend for os.Stderr.
	logFile := openFile("log.txt")
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, realLogFormat)
	backend1Formatter := logging.NewBackendFormatter(backend1, fileLogFormat)

	// Only errors and more severe messages should be sent to backend1
	// backend1Leveled := logging.AddModuleLevel(backend1)
	// backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend2Formatter, backend1Formatter)

	log.Debugf("debug %s", "secret")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}
