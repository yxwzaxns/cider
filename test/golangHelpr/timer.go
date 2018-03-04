package main

import (
	"time"
)

func main() {
	t := time.Now().Local().Add(1 * time.Second)
	println(t.String())
	time.Sleep(2 * time.Second)
	if t.After(time.Now().Local()) {
		println("ok")
		println(t.String())
	} else {
		println("false")
		println(t.String())
	}
}
