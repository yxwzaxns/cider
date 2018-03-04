package main

import (
	"time"
)

var queue chan int

func say(s string) {
	for i := 0; i < 5; i++ {
		queue <- i
	}
	time.Sleep(1500 * time.Millisecond)
	for i := 5; i < 10; i++ {
		queue <- i
	}
}

func recv() {
	for {
		// time.Sleep(500 * time.Millisecond)
		input := <-queue
		println(input)
	}
}

func compute() {
	// time.Sleep(5000 * time.Millisecond)
	queue <- 1 // when it completes, signal on the channel.
	println("ok")
	// queue <- 2
}

func main() {
	queue = make(chan int)
	go compute() // stat something in a goroutines
	time.Sleep(1500 * time.Millisecond)
	// println(len(queue))
	result := <-queue
	println(result)
	// queue <- 1
	// time.Sleep(1000 * time.Millisecond)
	// println(<-queue)
	// println("a")
	// for i := 0; i < 5; i++ {
	// 	queue <- i
	// }
	// go say("world")
	// time.Sleep(1500 * time.Millisecond)
	//
	// println(len(queue))
	// for {
	// 	if len(queue) != 0 {
	// 		println(<-queue)
	// 	}
	// 	//
	// }
	// go recv()
	// queue <- 3
	// time.Sleep(10000 * time.Millisecond)
}
