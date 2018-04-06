package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context, done chan int, c int) <-chan int {
	for {
		fmt.Printf("task %d running ...\n", c)
		time.Sleep(time.Duration(c) * time.Second)
	}
}

func Run(ctx context.Context, i int) {
	done := make(chan int)
	go run(ctx, done, i)
	select {
	case <-time.After(time.Duration(timeOut) * time.Second):
		fmt.Printf("task %d timeout!\n", i)
	case <-done:
		fmt.Printf("task %d was done.\n", i)
	}
}

func main() {
	ctx, shutDown := context.WithCancel(context.Background())
	go run(ctx, 10)
	time.Sleep(3 * time.Second)
	shutDown()
	time.Sleep(1 * time.Second)
}
