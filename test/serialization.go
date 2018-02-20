package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type ProjectStatus struct {
	Avtive        bool
	CiStatus      string // unstart working  finished
	CdStatus      string
	CurrentStatus string
}

type Project struct {
	ProjectName string
	ProjectURL  string
	ProjectStatus
}

func (p ProjectTable) FindByID(id int) []Project {
	var res []Project
	if id <= p.Size() {
		res = make([]Project, 1)
		res[0] = *p[id-1]
	}
	return res
}

func (p ProjectTable) FindAll() []Project {
	res := make([]Project, p.Size())
	for index := 0; index < p.Size(); index++ {
		res[index] = *p[index]
	}
	return res
}

func (p *ProjectTable) Add(project *Project) {
	*p = append(*p, project)
}

func (p ProjectTable) Size() int {
	return len(p)
}

type ProjectTable [](*Project)

func main() {
	projects := new(ProjectTable)
	project := Project{ProjectName: "blog", ProjectURL: "https://github.com/drone/drone-ui"}
	projects.Add(&project)
	println(projects.Size())

	file, err := os.Create("db.gob")
	if err != nil {
		fmt.Println(err)
	}
	enc := gob.NewEncoder(file)
	if err := enc.Encode(projects); err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
