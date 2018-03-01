package server

import (
	"cider/db"
	G "cider/global"
	"encoding/json"
	"fmt"
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

	payloadJsonString := c.PostForm("payload")

	var f interface{}

	if err := json.Unmarshal([]byte(payloadJsonString), &f); err != nil {
		panic(err)
	}

	var projectName string
	if firstParseJSON, ok := f.(map[string]interface{}); ok {
		pusher := firstParseJSON["repository"].(map[string]interface{})
		for k, v := range pusher {
			if k == "full_name" {
				projectName = v.(string)
			}
		}
	} else {
		panic(ok)
	}

	G.Core.AddTask("github.com/" + projectName)

	c.JSON(200, gin.H{
		"status": "ok",
	})
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

func dealProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	action := c.Param("action")
	switch action {
	case "stop":
		project := G.Projects.FindByID(id)
		G.Core.RemoveTask(project[0].ProjectURL)
		c.JSON(200, gin.H{
			"status": "ok",
			"info":   "stop req submit",
		})
		break
	case "submit":
		project := G.Projects.FindByID(id)
		G.Core.AddTask(project[0].ProjectURL)
		c.JSON(200, gin.H{
			"status": "ok",
			"info":   "submit req submit",
		})
		break
	default:
		c.JSON(200, gin.H{
			"status": "error",
			"reason": "unknown req : " + action,
		})
	}
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
		"status": "ok",
		"data":   projects,
	})
}

func createProject(c *gin.Context) {
	// url := c.PostForm("projectURL")
	var pr CreateProjectReq
	c.BindJSON(&pr)
	p := new(db.Project)
	p.ProjectName = "aong"
	p.ProjectURL = pr.ProjectURL
	G.Projects.Add(p)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func deleteProject(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "posted",
		"user":   "ok",
	})
}

func coreCheck(c *gin.Context) {
	item := c.Param("item")
	switch item {
	case "chan":
		G.EventsChan <- item
		break
	default:
		break
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func getTasks(c *gin.Context) {
	taskCount := G.Core.GetTaskCount()
	activeTaskCount := G.Core.GetActiveTaskCount()
	res := fmt.Sprintf("{'tasks count': %d,'active tasks count':%d}", taskCount, activeTaskCount)
	c.JSON(200, gin.H{
		"status":    "ok",
		"task info": res,
	})
}
