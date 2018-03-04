// package main
//
// import (
// 	"math/rand"
// 	"time"
// )
//
// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	a := rand.Intn(10)
// 	println("aaa", a)
// }

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(randString())
	}

	fmt.Println(b.String())
}

func randString() string {
	// Pretend to return a random string
	return "abc-123-"
}
