package server

import (
	"github.com/gin-gonic/gin"
)

//Core xx
type Core struct {
}

// Status xx
func (h Core) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

/*CiStart xxx
 */
func (h Core) CiStart() {

}

//CiDone xx
func (h Core) CiDone(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}

//CdDone xx
func (h Core) CdDone(c *gin.Context) {
	u := c.PostForm("username")

	c.JSON(200, gin.H{
		"status": "posted",
		"user":   u,
	})
}
