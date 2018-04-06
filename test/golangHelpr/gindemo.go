package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.Run() // listen and serve on 0.0.0.0:8080
	srv := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServeTLS("./certs/cider.aong.cn.cer", "./certs/cider.aong.cn.key"))

}
