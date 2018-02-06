package main

import (
	"regexp"
	"strings"
)

func convertURLToPath(url string) string {
	// func Replace(s, old, new string, n int) string
	return strings.Replace(url, "/", "_", -1)
}
func rege(s string) {
	res, _ := regexp.MatchString("too new", s)
	println(res)
}
func main() {
	rege("Error response from daemon: client version 1.36 is too new. Maximum supported API version is 1.35")
	// println(convertURLToPath("github.com/yxwzaxns/cider-ci-test"))
}
