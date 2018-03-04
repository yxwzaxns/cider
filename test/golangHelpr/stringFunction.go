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
func getString(s string) string {
	return s[0:17]
}

func split(s string, spe string) {
	println(len(strings.Split(s, spe)))
}

// var s = "adf"

func main() {
	// rege("Error response from daemon: client version 1.36 is too new. Maximum supported API version is 1.35")
	// println(getString("github.com/yxwzaxns/cider-ci-test"))
	split("github.com/yxwzaxns/cider", "/")
}
