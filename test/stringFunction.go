package main

import "strings"

func convertURLToPath(url string) string {
	// func Replace(s, old, new string, n int) string
	return strings.Replace(url, "/", "_", -1)
}
func main() {
	println(convertURLToPath("github.com/yxwzaxns/cider-ci-test"))
}
