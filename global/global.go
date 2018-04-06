package global

import (
	"github.com/yxwzaxns/cider/types"
)

var (
	SysChan  chan types.CR
	StopChan chan interface{}
)

func Init() {
	SysChan = make(chan types.CR, 10)
	StopChan = make(chan interface{})
}
