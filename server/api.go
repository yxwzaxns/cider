package server

import (
	"github.com/gin-gonic/gin"
)

//User xx
type User struct {
	Name string `json:"name"`
}

//Login xx
func (h User) Login(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}

//Logout xx
func (h User) Logout(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}
