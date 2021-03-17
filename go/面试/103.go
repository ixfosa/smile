package main

import (
	"encoding/json"
	"fmt"
)

//下面这段代码输出什么？
type People103 struct {
	name string `json:"name"`
}
func main() {
	jsonStr := `{
		"name" : "ixfosa"
	}`

	var p People103

	err := json.Unmarshal([]byte(jsonStr), &p)

	if err != nil {
		panic(err)
	}
	fmt.Println(p) // {} 结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体。
}
