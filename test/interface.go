package main

import "encoding/json"

type Any interface {
}

func t() (v int, err bool) {
	v = 3
	err = true
	return
}
func main() {
	b := []byte(`{"Title":"Go语言编程","Authors":"ljnladusf","Publisher":"ituring.com.cn","IsPublished":"sdf","Price":"fklmdf","Sales":"dff"}`)

	var f interface{}

	if err := json.Unmarshal([]byte(b), &f); err != nil {
		panic(err)
	}

	if payloadJson, ok := f.(map[string]interface{}); ok {
		// fmt.Println(reflect.TypeOf(payloadJson["pusher"]).Field(0))
		// fmt.Println(reflect.TypeOf(payloadJson))
		// for k, v := range payloadJson {
		// 	println(k, ":", v)
		// }
		for _, v := range payloadJson {
			println(v.(string))
		}
	} else {
		panic(ok)
	}
}
