package server

import (
	"github.com/gin-gonic/gin"
)

//NewRouter xx
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/assets", "./assets")
	// router.LoadHTMLGlob(ProjectPath + "templates/*")

	// health := new(controllers.HealthController)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/")
	// router.Use(AuthMiddleware())
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(User)
			userGroup.POST("/login", user.Login)
			userGroup.POST("/logout", user.Logout)
		}
	}

	return router
}
