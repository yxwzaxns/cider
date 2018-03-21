package main

import "github.com/davecgh/go-spew/spew"

type Project struct {
	ProjectName string
	AutoBuild   bool
	AutoDeploy  bool

	CINotification bool
	CDNotification bool
	Email          string

	PauseServer bool
}

func (p *Project) Update(field string, value interface{}) {
	switch field {
	case "AutoBuild":
		p.AutoBuild = value.(bool)
		break
	case "AutoDeploy":
		p.AutoDeploy = value.(bool)
		break
	case "CINotification":
		p.CINotification = value.(bool)
		break
	case "CDNotification":
		p.CDNotification = value.(bool)
		break
	case "PauseServer":
		p.PauseServer = value.(bool)
		break
	case "Email":
		p.Email = value.(string)
		break
	default:
		break
	}
}

type ProjectTable []*Project

func (p *ProjectTable) FindByName(projectName string) *Project {
	var res *Project
	if p.Size() > 0 {
		for _, _p := range *p {
			if _p.ProjectName == projectName {
				// res = make([]Project, 1)
				res = _p
			}
		}
		return res
	}
	return nil
}

func (p *ProjectTable) Get(projectName string) *Project {
	if p.Size() > 0 {
		for _, _p := range *p {
			if _p.ProjectName == projectName {
				// res = make([]Project, 1)
				return _p
			}
		}
	}
	return nil
}

func (p *ProjectTable) Add(project *Project) {
	*p = append(*p, project)
}

func (p *ProjectTable) Size() int {
	return len(*p)
}

var Projects ProjectTable

func main() {
	p := Project{ProjectName: "aong1"}
	p1 := Project{ProjectName: "aong2"}
	p2 := Project{ProjectName: "aong3"}
	Projects.Add(&p)
	Projects.Add(&p1)
	Projects.Add(&p2)
	spew.Dump(Projects.Get("aong2"))
	Projects.Get("aong2").Update("Email", "adfadsf@aong.cn")
	Projects.Get("aong2").Update("AutoBuild", true)
	spew.Dump(Projects.Get("aong2"))

	// println(reflect.TypeOf(Projects[0]).Field(1).Name)
}
