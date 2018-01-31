package db

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
