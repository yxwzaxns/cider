package main

// "os",
import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
}
