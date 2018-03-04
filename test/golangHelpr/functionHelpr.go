package main

import "fmt"

type Project struct {
	Name string
}

func test(projectName string) interface{} {
	if projectName == "abc" {
		return nil
	} else if projectName == "abc1" {
		return []Project{Project{Name: projectName}, Project{Name: projectName}}
	} else {
		return Project{Name: projectName}
	}
}

func main() {
	name := "abc1"
	res := test(name)
	fmt.Printf("%T: %v", res, res)
}
