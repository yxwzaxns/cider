package db

import (
	"testing"

	"github.com/yxwzaxns/cider/utils"
)

var projectsTestSample [4]string
var tmpDBPath string

func TestCreateDB(t *testing.T) {
	t.Log("Start Do Create DB Test")
	tmpDBPath = "./cider.test.db"
	Init(tmpDBPath)
}

func TestInsertData(t *testing.T) {
	projectsTestSample = [4]string{"github.com/yxwzaxns/cider",
		"github.com/yxwzaxns/cider-ui",
		"github.com/yxwzaxns/cider-ci-test",
		"github.com/yxwzaxns/cider-client"}
	for _, p := range projectsTestSample {
		Projects.Create(p)
	}
}

func TestUpdateData(t *testing.T) {
	for _, p := range Projects {
		p.Update("Email", "test@example.com")
	}

	for _, p := range Projects {
		if p.Email != "test@example.com" {
			t.Fatal("Update test failed")
		} else {
			t.Log("Update test succeed")
		}
	}
}

func TestQueryData(t *testing.T) {
}

func TestDeleteData(t *testing.T) {
	if !Projects.Get(projectsTestSample[0]).Delete() {
		t.Fatal("Delete test failed")
	} else {
		println("Delete test succeed")
	}
}

func TestSaveDB(t *testing.T) {
	SaveDb()
}

func TestRecoveryDB(t *testing.T) {
	RebuildDb()
	if Projects.Size() == 3 {
		println("Rebuild DB succeed")
	} else {
		t.Fatal("Rebuild DB Failed")
	}

}

func TestDropDB(t *testing.T) {
	utils.DeleteFile(tmpDBPath)
}
