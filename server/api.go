package server

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
}

func (h User) Login(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}

func (h User) Logout(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}
