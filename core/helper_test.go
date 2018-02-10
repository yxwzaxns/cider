package core

import (
	"testing"
)

func TestParsePort(t *testing.T) {
	p := "127.0.0.1:8000:80/tcp"
	pt := [2][2]string{[2]string{"127.0.0.1", "8000"}, [2]string{"80/tcp", ""}}
	res := ParsePort(p)
	if pt != res {
		t.Error()
	}
}
