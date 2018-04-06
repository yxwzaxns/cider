package main

import (
	"context"
	"fmt"
	"time"
)

var tasks = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var timeOut int = 5

func run(ctx context.Context, i int) {
	select {
    case ctx.
  }
}
func Run(ctx context.Context, i int) {
	go run(ctx, i)
	select {
	case <-time.After(5 * time.Second):
		ctx.Done()
	}
}
func main() {
	ctx := context.Background()
	for i := 0; i < len(tasks); i++ {
		go Run(ctx, i)
	}
	time.Sleep(12 * time.Second)
}
