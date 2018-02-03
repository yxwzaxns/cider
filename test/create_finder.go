package main

import (
	"os"
	"path/filepath"
)

func main() {
	basePath, _ := filepath.Abs("/")
	path := filepath.Join(basePath, "tmp", "cider", "github.com/yxwzaxns/aong")
	if _, err := os.Stat(path); os.IsExist(err) {
		os.RemoveAll(path)
	}
	if err := os.Mkdir(path, 0777); err != nil {
		println(err.Error())
	}

}
