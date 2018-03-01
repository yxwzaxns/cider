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
				})
			})
			// about project
			projectGroup := v1.Group("project")
			{
				projectGroup.GET("/:id", getProject)
				// projectGroup.OPTIONS("/:id", preflight)
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
	}
	return router
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		// c.Set("example", "12345")

		// before request
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			// c.Header("Content-Type", "application/x-www-form-urlencoded")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Authorization, Content-Type")
			c.JSON(204, gin.H{})
		}

		if c.Request.Header.Get("Origin") != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Next()
		} else {
			c.Next()
		}

		// c.Next()

		// after request
		// latency := time.Since(t)
		// log.Print(latency)

		// access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)
	}
}
