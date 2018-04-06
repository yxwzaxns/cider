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

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	a[0] = 2333
	return min
}

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
