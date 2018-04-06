package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yxwzaxns/cider/config"
	"github.com/yxwzaxns/cider/core"
	"github.com/yxwzaxns/cider/db"
	"github.com/yxwzaxns/cider/global"
	"github.com/yxwzaxns/cider/server/middlewares"
)

// Init xx
func Init() {
	middleware.Init()
	core.Init(global.SysChan)
	// rebuild database from db file if need to do
	db.Init(config.Conf.AppDbPath)
	db.RebuildDb()

	r := NewRouter()
	srv := &http.Server{
		Addr:    config.Conf.ListenIP + `:` + config.Conf.ListenPort,
		Handler: r,
	}

	go func() {
		// service connections
		// if err := srv.ListenAndServe(); err != nil {
		// 	G.Log.Debugf("listen: %s", err)
		// }
		// G.SysChan <- fmt.Sprintf("Listen At %s:%s", G.Config.ListenIP, G.Config.ListenPort)
		if err := srv.ListenAndServeTLS("certs/cider.aong.cn.cer", "certs/cider.aong.cn.key"); err != nil {
			config.Log.Debugf("Listen error: %s !!!!", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	config.Log.Info("Shutdown Server ...")
	global.StopChan <- 1
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		config.Log.Fatal("Server Shutdown:", err)
	}
	finishWork()
	config.Log.Info("Server exiting")
}

func finishWork() {
	config.Log.Debug("Save database into file")
	db.SaveDb()
}
