package db

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

type ProjectDB [](*Project)
