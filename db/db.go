package db

var dbPath string

type ProjectStatus struct {
	Avtive        bool
	CiStatus      string
	CdStatus      string
	CurrentStatus string
}

type Project struct {
	ProjectName string
	ProjectURL  string
	CreatedTime string

	AutoBuild  bool
	AutoDeploy bool

	CINotification bool
	CDNotification bool
	Email          string

	PauseServer bool

	ProjectStatus
}

type ProjectTable [](*Project)
