package main

import (
	"fmt"
)

// 下面代码有什么问题？
func main() {
	// var str string = nil // cannot use nil as type string in assignment
	// 修复
	var str string
	if str == "" {
		fmt.Println("nil")
	}
}
