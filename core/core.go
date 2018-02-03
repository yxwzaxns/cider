package core

import (
	"time"
)

type M struct {
	URL  string
	info string
}

type Core struct {
	Tasks tasks
}

var (
	eventsChan chan string
	ciChan     chan M
	cdChan     chan M
)

func (c *Core) GetTaskCount() int {
	return c.Tasks.Size()
}
func (c *Core) GetActiveTaskCount() int {
	return c.Tasks.ActiveTaskCount()
}
func (c *Core) createEventsDaemon() {
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			if len(eventsChan) != 0 {
				c.dealEvent(<-eventsChan)
			}
		}
	}()
}

func (c *Core) dealEvent(event string) {
	println("<-------------", event, "--------------->")
}

func (c *Core) Init(EventChan chan string) {
	eventsChan = EventChan
	ciChan = make(chan M, 10)
	cdChan = make(chan M, 10)
	c.createEventsDaemon()
}

func (c *Core) RemoveTask(project string) {
	_, err := c.Tasks.FindTaskByURL(project)
	if err != "" {
		// delete task form tasks
		// to do
	}
}
func (c *Core) AddTask(project string) {
	// check the project if exist a task
	// create a new task if the task not exist
	_, err := c.Tasks.FindTaskByURL(project)
	if err != "" {

	} else {
		c.Tasks.CreateTask(project)
	}
	c.StartTask(project)
}

func (c *Core) StartTask(project string) {
	task, err := c.Tasks.FindTaskByURL(project)
	if err == "" {
		go task.Run()
		eventsChan <- "start run task"
	} else {
		eventsChan <- err
	}
}
