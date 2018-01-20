package server

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
}

func (h User) Signup(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}
