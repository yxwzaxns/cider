package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		// c.Set("example", "12345")

		// before request
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			// c.Header("Content-Type", "application/x-www-form-urlencoded")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Authorization, Content-Type")
			c.AbortWithStatusJSON(204, gin.H{})
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
