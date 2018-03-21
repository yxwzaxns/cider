package db

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
