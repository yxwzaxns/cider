package core

type Task struct {
	Active bool
	Status string //unstart,ready,pending,finished
	URL    string
}

type tasks [](*Task)

func (t *Task) Run() {
	t.Ready()
	t.CI()
	t.CD()
	t.Done()
}
func (t *Task) Ready() {
	t.Active = true
	// fmt.Printf("goroutine :  %p", t)
	// randTime := rand.Intn(10000)
}
func (t *Task) CI() {
	t.Status = "CI"
	StartCI(t.URL, EventChan)
}
func (t *Task) CD() {
	t.Status = "CD"
	StartCD(t.URL, EventChan)
}
func (t *Task) Done() {
	t.Active = false
	t.Status = "finished"
	res := t.URL + "  finished"

	// M := new(ciderTypes.CR)
	M.ProjectName = t.URL
	M.TaskStage = "DONE"
	M.Message = res
	EventChan <- M
}

func (t *tasks) CreateTask(project string) {
	task := new(Task)
	task.Active = false
	task.Status = "ready"
	task.URL = project
	*t = append(*t, task)
}

func (t *tasks) Size() int {
	return len(*t)
}

func (t *tasks) ActiveTaskCount() int {
	count := 0
	for index := 0; index < t.Size(); index++ {
		if (*t)[index].Active != false {
			count++
		}
	}
	return count
}
func (t *tasks) FindTaskByURL(url string) (*Task, string) {
	var aimTask *Task
	err := ""
	for index := 0; index < t.Size(); index++ {
		if (*t)[index].URL == url {
			aimTask = (*t)[index]
		} else {
			err = "not find task"
		}
	}
	return aimTask, err
}
