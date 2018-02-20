package main

type Any interface {
}

func t() (v int, err bool) {
	v = 3
	err = true
	return
}
func main() {
	// var a Any
	// a = 5

	// v, err := a.(int)
	// v := t()
	// // if err {
	// println(v)
	// }
}
