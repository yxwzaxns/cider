package db

var dbPath string

type ProjectStatus struct {
	Active        bool
	CiStatus      string
	CdStatus      string
	CurrentStatus string
}

type Log struct {
	LastTime string
	LastLog  string
}

type ProjectSetting struct {
	AutoBuild  bool
	AutoDeploy bool

	CINotification bool
	CDNotification bool
	Email          string

	PauseServer bool
}
type Project struct {
	ProjectName string
	ProjectURL  string
	CreatedTime string

	ProjectSetting
	ProjectStatus
	Log
}

type ProjectTable [](*Project)
