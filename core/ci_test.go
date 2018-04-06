package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/yxwzaxns/cider/types"
)

func TestStartCI(t *testing.T) {
	mChan := make(chan types.CR, 10)
	codeURL := "github.com/yxwzaxns/aong-ghost"
	println("start test ci")
	go func(mChan chan types.CR) {
		for {
			time.Sleep(500 * time.Millisecond)
			if len(mChan) != 0 {
				m := <-mChan
				fmt.Println("<-------------", m.Message, "--------------->")
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
