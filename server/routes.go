package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewRouter xx
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CorsMiddleware())
	router.Use(AuthMiddleware())
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
	Api := router.Group("api")
	{
		v1 := Api.Group("v1")
		{
			//
			v1.GET("ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "API Server Working",
					"status":  "ok",
				})
			})
			// about project
			projectGroup := v1.Group("project")
			{
				projectGroup.GET("/:name", getProject)
				// projectGroup.OPTIONS("/:id", preflight)
				projectGroup.GET("/:name/:action", dealProject)
				projectGroup.POST("/", createProject)
				projectGroup.PUT("/:name/update", updateProject)
				projectGroup.DELETE("/", deleteProject)
			}
			// about user
			userGroup := v1.Group("user")
			{
				userGroup.POST("/auth", Auth)
				userGroup.POST("/dissauth", DissAuth)
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
	}
	return router
}
