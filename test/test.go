package main

import "github.com/davecgh/go-spew/spew"

type ProjectStatus struct {
	Avtive   bool
	CiStatus bool
	CdStatus bool
	Now      string
}

type Project struct {
	ProjectName string
	ProjectURL  string
	ProjectStatus
}

type Pa [](*Project)

func (p Pa) Get(id int) Project {
	return *p[id]
}

func (p *Pa) Set(project *Project) {
	*p = append(*p, project)
}

var Projects Pa

func main() {
	p := new(Project)
	p.ProjectName = "aong"
	p.ProjectURL = "https://ssdfdf"

	Projects.Set(p)

	spew.Dump(Projects.Get(0))

	p = new(Project)
	p.ProjectName = "aaaa"
	p.ProjectURL = "https://ssvvvvf"
	// p.Status = new(ProjectStatus)
	Projects.Set(p)

	spew.Dump(Projects.Get(1))

	println(cap(Projects))
	println(len(Projects))
}
