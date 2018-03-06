package server

import (
	"cider/db"
	G "cider/global"
	"cider/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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

	if G.Projects.Has(projectName) != false {
		G.Core.AddTask("github.com/" + projectName)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(403, gin.H{
			"status": "not find project",
		})
	}

}

//GitlabHook xx
func GitlabHook(c *gin.Context) {

}

//Login xx
func Auth(c *gin.Context) {
	var ar AuthReq
	c.BindJSON(&ar)
	if ar.Key != utils.GetKey() {
		c.JSON(200, gin.H{
			"code":   401,
			"status": "key not match",
		})
	} else {
		token := NewToken()

		c.JSON(200, gin.H{
			"status": "ok",
			"token":  token,
		})
	}
}

//Logout xx
func DissAuth(c *gin.Context) {
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
	// ProjectURL demo : yxwzaxns/cider
	var pr CreateProjectReq
	c.BindJSON(&pr)
	projectInfo := strings.Split(pr.ProjectURL, "/")
	if len(projectInfo) != 3 {
		c.JSON(200, gin.H{
			"status": "error projrct url",
		})
	}
	p := new(db.Project)
	p.ProjectName = projectInfo[2]
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
