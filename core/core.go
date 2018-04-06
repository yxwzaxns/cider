package core

import (
	"github.com/yxwzaxns/cider/config"
	"github.com/yxwzaxns/cider/global"
	"github.com/yxwzaxns/cider/types"
)

var (
	Core      *core
	EventChan chan types.CR
)

type core struct {
	Tasks tasks
}

func Init(sysChan chan types.CR) {
	Core = new(core)
	EventChan = sysChan
	Core.createEventsDaemon()
}

func (c *core) GetTaskCount() int {
	return c.Tasks.Size()
}
func (c *core) GetActiveTaskCount() int {
	return c.Tasks.ActiveTaskCount()
}
func (c *core) createEventsDaemon() {
	go func() {
		for {
			select {
			case <-global.StopChan:
				config.Log.Info("Core Events Daemon Ready To Close")
				return
			case cr := <-EventChan:
				c.dealEvent(cr)
			}
		}
	}()
}

func (c *core) dealEvent(message types.CR) {
	config.Log.Debugf("<------------- %s --------------->", message.Message)
}

func (c *core) RemoveTask(project string) {
	_, err := c.Tasks.FindTaskByURL(project)
	if err != "" {
		// delete task form tasks
		// to do
	}
}
func (c *core) AddTask(project string) {
	// check the project if exist a task
	// create a new task if the task not exist
	_, err := c.Tasks.FindTaskByURL(project)
	if err != "" {

	} else {
		c.Tasks.CreateTask(project)
	}
	c.StartTask(project)
}

func (c *core) StartTask(project string) {
	task, err := c.Tasks.FindTaskByURL(project)
	if err == "" {
		go task.Run()
		// EventChan <- "start run task"
	} else {
		// EventChan <- err
	}
}
