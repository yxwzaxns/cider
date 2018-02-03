package core

import (
	"testing"
)

func TestStartCI(t *testing.T) {
	mChan := make(chan M, 10)
	codeURL := "github.com/yxwzaxns/aong-ghost"
	println("start test ci")
	StartCI(codeURL, mChan)

}

func TestCreateTarFile(t *testing.T) {
	TestPath := "/tmp/cider_workspace/"
	createTar(TestPath)
}
func TestDeleteFile(t *testing.T) {
	TestPath := "/tmp/cider_workspace/aim.tar"
	deleteFile(TestPath)
}
