package server

import (
	G "cider/global"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//NewRouter xx
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/assets", path.Join(G.BasePath, "server/assets"))
	router.LoadHTMLGlob(path.Join(G.BasePath, "server/templates/*"))

	// health := new(controllers.HealthController)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	})

	dashboard := router.Group("dashboard")
	{
		dashboard.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "cider system dashboard",
			})
		})
	}
	// router.Use(AuthMiddleware())
	v1 := router.Group("v1")
	{
		//
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "API Server Working",
			})
		})
		// about project
		projectGroup := v1.Group("project")
		{
			projectGroup.GET("/:id", getProject)
			projectGroup.GET("/:id/:action", dealProject)
			projectGroup.POST("/", createProject)
			projectGroup.DELETE("/", deleteProject)
		}
		// about user
		userGroup := v1.Group("user")
		{
			user := new(User)
			userGroup.POST("/login", user.Login)
			userGroup.POST("/logout", user.Logout)
		}
		// about hook
		hookGroup := v1.Group("hook")
		{
			hookGroup.POST("/github", GithubHook)
			hookGroup.POST("/gitlab", GitlabHook)
		}
		// about core
		coreGroup := v1.Group("core")
		{
			// core := new(Core)
			coreGroup.GET("/check/:item", coreCheck)
			coreGroup.GET("/task", getTasks)
			// coreGroup.POST("/ci-done", core.CiDone)
			// coreGroup.POST("/cd-done", core.CdDone)
		}
	}

	return router
}
