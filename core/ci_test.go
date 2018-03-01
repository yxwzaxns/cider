package core

import (
	"fmt"
	"testing"
	"time"
)

func TestStartCI(t *testing.T) {
	mChan := make(chan M, 10)
	codeURL := "github.com/yxwzaxns/aong-ghost"
	println("start test ci")
	go func(mChan chan M) {
		for {
			time.Sleep(500 * time.Millisecond)
			if len(mChan) != 0 {
				m := <-mChan
				fmt.Println("<-------------", m.info, "--------------->")
			}
		}
	}(mChan)
	StartCI(codeURL, mChan)
	time.Sleep(5000 * time.Millisecond)

}

func TestCreateTarFile(t *testing.T) {
	TestPath := "/tmp/cider_workspace/"
	createTar(TestPath)
}
func TestDeleteFile(t *testing.T) {
	TestPath := "/tmp/cider_workspace/aim.tar"
	deleteFile(TestPath)
}
