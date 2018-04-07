package db

import (
	"errors"
	"strings"

	"github.com/yxwzaxns/cider/utils"
)

func (p *ProjectTable) FindByID(id int) []Project {
	var res []Project
	if id <= p.Size() {
		res = make([]Project, 1)
		res[0] = *(*p)[id-1]
	}
	return res
}

func (p *ProjectTable) FindByName(projectName string) interface{} {
	var res []Project
	if p.Size() > 0 {
		for _, _p := range *p {
			if _p.ProjectName == projectName {
				res = make([]Project, 1)
				res[0] = *_p
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
				return _p
			}
		}
	}
	return nil
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

func (p *Project) Delete() bool {
	for i, _p := range Projects {
		if _p.ProjectName == p.ProjectName {
			Projects = append(Projects[:i], Projects[i+1:]...)
		}
	}
	if !Projects.Has(p.ProjectName) {
		return true
	}
	return false
}

func (p *ProjectTable) Has(projectName string) bool {
	if 0 < p.Size() {
		for _, _p := range *p {
			if _p.ProjectName == projectName {
				return true
			}
		}
	}
	return false
}

func (p *ProjectTable) Create(url string) error {
	// detecting whether a project is exists or not.
	projectInfo := strings.Split(url, "/")
	if p.Has(projectInfo[2]) {
		return errors.New("Project has exists")
	}
	// create a new project.
	np := new(Project)
	np.ProjectName = projectInfo[2]
	np.ProjectURL = url
	np.CreatedTime = utils.GetCurrentTime()

	// np.ProjectStatus
	np.Active = false
	np.CdStatus = ""
	np.CiStatus = ""
	np.CurrentStatus = ""

	// np.ProjectSetting
	np.AutoBuild = true
	np.AutoDeploy = true
	np.CDNotification = false
	np.CINotification = false
	np.PauseServer = false
	np.Email = ""

	// np.Log
	np.Log.LastLog = ""
	np.Log.LastTime = ""

	p.Add(np)
	return nil
}
func (p *ProjectTable) FindAll() []Project {
	res := make([]Project, p.Size())
	for index := 0; index < p.Size(); index++ {
		res[index] = *(*p)[index]
	}
	return res
}

func (p *ProjectTable) Add(project *Project) {
	*p = append(*p, project)
}

func (p *ProjectTable) Size() int {
	return len(*p)
}
