package main

import "fmt"

func main() {
	var data interface{} = "great"

	// data 混用
	if data, ok := data.(int); ok {
		fmt.Println("[is an int], data: ", data)
	} else {
		fmt.Println("[not an int], data: ", data) // [isn't a int], data:  0
	}
}
