package server

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/yxwzaxns/cider/utils"

	"github.com/gin-gonic/gin"
	"github.com/yxwzaxns/cider/config"
	"github.com/yxwzaxns/cider/core"
	"github.com/yxwzaxns/cider/db"
	"github.com/yxwzaxns/cider/server/middlewares"
)

// web hook

//GithubHook xx
func GithubHook(c *gin.Context) {
	// check Auth

	payloadJSONString := c.PostForm("payload")

	var f interface{}

	if err := json.Unmarshal([]byte(payloadJSONString), &f); err != nil {
		panic(err)
	}

	var projectName string
	if firstParseJSON, ok := f.(map[string]interface{}); ok {
		pusher := firstParseJSON["repository"].(map[string]interface{})
		for k, v := range pusher {
			if k == "name" {
				projectName = v.(string)
			}
		}
	} else {
		panic(ok)
	}

	if project := db.Projects.Get(projectName); project != nil {
		core.Core.AddTask(project.ProjectURL)
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

//TestHook is test of system
func TestHook(c *gin.Context) {
	config.Log.Debug("start test system")
	core.Core.AddTask("github.com/yxwzaxns/cider-ci-test")
}

//Auth xx
func Auth(c *gin.Context) {
	var ar AuthReq
	c.BindJSON(&ar)
	if ar.Key != utils.GetKey() {
		c.JSON(200, gin.H{
			"code":   401,
			"status": "key not match",
		})
	} else {
		token := middleware.NewToken()

		c.JSON(200, gin.H{
			"status": "ok",
			"token":  token,
		})
	}
}

//DissAuth xx
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
		project := db.Projects.FindByID(id)
		core.Core.RemoveTask(project[0].ProjectURL)
		c.JSON(200, gin.H{
			"status": "ok",
			"info":   "stop req submit",
		})
		break
	case "submit":
		project := db.Projects.FindByID(id)
		core.Core.AddTask(project[0].ProjectURL)
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
func updateProject(c *gin.Context) {
	project := db.Projects.Get(c.Param("name"))
	var projectItem UpdateProjectItem
	c.BindJSON(&projectItem)
	var field, value = utils.ParseField(projectItem.Field)
	db.Projects.Get(project.ProjectName).Update(field, value)

	c.JSON(200, gin.H{
		"status": db.Projects.Get(project.ProjectName),
	})
}
func getAllProject(c *gin.Context) {
	projects := db.Projects.FindAll()
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   projects,
	})
}

func getProject(c *gin.Context) {
	n := c.Param("name")
	if n == "all" {
		getAllProject(c)
	} else {
		projects := db.Projects.Get(n)
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   projects,
		})
	}
}

func createProject(c *gin.Context) {
	// ProjectURL demo : github.com/yxwzaxns/cider
	var pr CreateProjectReq
	c.BindJSON(&pr)
	projectInfo := strings.Split(pr.ProjectURL, "/")
	if len(projectInfo) != 3 {
		c.JSON(200, gin.H{
			"status": "error projrct url",
		})
	}
	if err := db.Projects.Create(pr.ProjectURL); err != nil {
		c.JSON(200, gin.H{
			"status": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}

}

func deleteProject(c *gin.Context) {
	name := c.Param("name")
	println(name)
	if db.Projects.Get(name).Delete() {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "failed",
		})
	}
}

func coreCheck(c *gin.Context) {
	item := c.Param("item")
	switch item {
	case "chan":
		// EventsChan <- item
		break
	default:
		break
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func getTasks(c *gin.Context) {
	taskCount := core.Core.GetTaskCount()
	activeTaskCount := core.Core.GetActiveTaskCount()
	res := fmt.Sprintf("{'tasks count': %d,'active tasks count':%d}", taskCount, activeTaskCount)
	c.JSON(200, gin.H{
		"status":    "ok",
		"task info": res,
	})
}
