package main

import (
	"io/ioutil"
	"path/filepath"
)

func main() {
	files := getFiles("/tmp/cider_workspace/github_com_yxwzaxns_cider-ci-test/")
	for _, n := range files {
		println(n)
	}
}

func getDirList(path string) []string {
	files := []string{}
	fs, _ := ioutil.ReadDir(path)
	for _, f := range fs {
		files = append(files, filepath.Join(path, f.Name()))
	}
	return files
}
