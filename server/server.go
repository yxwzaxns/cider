package server

import (
	G "cider/global"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Init xx
func Init() {
	// init CI/CD queue
	// CDQUEUE := make(chan string)
	// c := new(config.Config)
	// c.Init()
	// rebuild database from db file
	G.Projects.RebuildDb()

	r := NewRouter()
	srv := &http.Server{
		Addr:    G.Config.ListenIP + `:` + G.Config.ListenPort,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			G.Log.Debugf("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	G.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		G.Log.Fatal("Server Shutdown:", err)
	}
	finishWork()
	G.Log.Info("Server exiting")
}

func finishWork() {
	G.Log.Debug("Save database into file")
	G.Projects.SaveDb()
}
