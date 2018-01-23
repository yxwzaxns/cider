package main

import "path/filepath"

func main() {
	path, _ := filepath.Abs("../")
	print(path)
}
