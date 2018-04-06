package main

import (
	"fmt"
	"time"
)

var queue chan interface{}

func log(s interface{}) {
	fmt.Println("sys real log : %s", s)
}
func sysChan(done chan int) {
	for {
		// time.Sleep(500 * time.Millisecond)
		// input := <-queue
		// println(input)
		select {
		case s := <-queue:
			log(s)
		case <-done:
			println("syschan finished")
			break
		}
	}
}

func main() {
	queue = make(chan interface{}, 10)
	done := make(chan int)
	// go compute() // stat something in a goroutines
	go sysChan(done)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		queue <- i
	}

	time.Sleep(11 * time.Second)
	done <- 1
	// println(len(queue))
	// result := <-queue
	// println(result)

}
