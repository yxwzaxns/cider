package server

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude auth url
		if c.Request.RequestURI == "/api/v1/user/auth" || c.Request.RequestURI == "/api/v1/ping" {
			c.Next()
			return
		}
		//
		println("token: ", c.Request.Header.Get("Authorization"))
		if c.Request.Header.Get("Authorization") != "" {
			if permit := CheckPermit(c.Request.Header.Get("Authorization")); !permit {
				// c.JSON(403, gin.H{})
				c.AbortWithStatus(403)
				// return
			} else {
				UpdateTokenExpireTime()
				c.Next()
				return
			}
		} else {
			// c.JSON(403, gin.H{})
			// return
			c.AbortWithStatus(403)
		}

	}
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
			return
		}

		if c.Request.Header.Get("Origin") != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Next()
		} else {
			c.Next()
		}
		return
		// c.Next()

		// after request
		// latency := time.Since(t)
		// log.Print(latency)

		// access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)
	}
}
