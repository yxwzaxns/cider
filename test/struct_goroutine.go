package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var tchan chan int

type Key struct {
	Active bool
	Value  int
}

type Keys [](*Key)
type Demo struct {
	Arr Keys
}

var D Demo

func (k *Keys) Add(key Key) {
	*k = append(*k, &key)
}

func (t *Keys) Size() int {
	return len(*t)
}
func (t *Keys) Asize() int {
	count := 0
	for index := 0; index < t.Size(); index++ {
		if (*t)[index].Active != false {
			count++
		}
	}
	return count
}
func (k *Key) Test() {
	k.Value = 2
	k.Active = true
	spew.Dump(k)
	time.Sleep(3000 * time.Millisecond)
	fmt.Printf(" self Arr address is : %p \n", k)
	tchan <- 1
	// fmt.Printf(" out Arr address is : %p \n", key)
	println("goroutine finished")
	tchan <- 1
}
func main() {
	tchan = make(chan int)
	fmt.Printf("Arr len is: %d, Arr address is : %p \n", len(D.Arr), &D.Arr)
	k := new(Key)
	k.Value = 1
	k.Active = false
	D.Arr.Add(*k)
	spew.Dump(D.Arr[0])
	fmt.Printf("Arr len is: %d, Arr address is : %p \n", len(D.Arr), D.Arr[0])
	go D.Arr[0].Test()
	<-tchan
	println("active count is : ", D.Arr.Asize())
	<-tchan
}
