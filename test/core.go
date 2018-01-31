package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Task struct {
	Status bool
	URL    string
}

type tasks []Task

type Core struct {
	Tasks tasks
}

func (t Task) Run() {
	if false {
		defer func() {
			println("Run Task")
		}()
	}
	randTime := rand.Intn(10000)
	time.Sleep(time.Duration(randTime) * time.Millisecond)
	// println(t.URL, "use", randTime)
	res := t.URL + strconv.Itoa(randTime)
	EventsChan <- res

	t.Init()
	t.CI()
	t.CD()
	t.Done()
}
func (t Task) Init() {

}
func (t Task) CI() {

}
func (t Task) CD() {

}
func (t Task) Done() {

}

func (t *tasks) CreateTask(project string) {
	task := new(Task)
	task.Status = false
	task.URL = project
	*t = append(*t, *task)
}

func (t tasks) Size() int {
	return len(t)
}

func (t tasks) FindTaskByURL(url string) Task {
	aimTask := new(Task)
	for index := 0; index < t.Size(); index++ {
		if t[index].URL == url {
			aimTask = &t[index]
		}
	}
	return *aimTask
}

func (c Core) Init() {

}
func (c Core) addTask(project string) {
	defer func() {
		println("addTask")
	}()
	c.Tasks.CreateTask(project)
	c.StartTask(project)
}

func (c Core) StartTask(project string) {
	defer func() {
		println("createTask")
	}()
	task := c.Tasks.FindTaskByURL(project)
	go task.Run()
}

func dealEvents() {
	// for e := range EventsChan {
	// 	println(e)
	// }
	for {
		if len(EventsChan) != 0 {
			input := <-EventsChan
			fmt.Printf("%s ", input)
		}
	}
}

var EventsChan = make(chan string, 10)

func main() {
	rand.Seed(time.Now().UnixNano())
	core := new(Core)
	core.Init()
	for index := 0; index < 10; index++ {
		project := "yxwzaxns/cider" + strconv.Itoa(index)
		core.addTask(project)
	}
	go dealEvents()
	time.Sleep(11000 * time.Millisecond)

}
