package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // cancel when we are finished consuming integers
	a := make(chan int)
	go func() {
		for {
			a <- 1
		}
	}()
	for n := range a {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
