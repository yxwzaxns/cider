package server

import (
	"cider/db"
	G "cider/global"
	"strconv"

	"github.com/gin-gonic/gin"
)

//User xx
type User struct {
	Name string `json:"name"`
}

// web hook

//GithubHook xx
func GithubHook(c *gin.Context) {
	// check Auth

	// call start Continuous integration

}

//GitlabHook xx
func GitlabHook(c *gin.Context) {

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

func getAllProject() []db.Project {
	projects := G.Projects.FindAll()
	return projects
}

func getProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var projects []db.Project
	if id != 0 {
		projects = G.Projects.FindByID(id)
	} else {
		projects = getAllProject()
	}
	c.JSON(200, gin.H{
		"status":   "ok",
		"projects": projects,
	})
}

func createProject(c *gin.Context) {
	url := c.PostForm("url")
	p := new(db.Project)
	p.ProjectName = "aong"
	p.ProjectURL = url
	G.Projects.Add(p)

	c.JSON(200, gin.H{
		"status": "ok",
		"url":    url,
	})
}

func deleteProject(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "posted",
		"user":   "ok",
	})
}
