package main

import (
	"fmt"
)

// 下面代码输出什么？
func main() {
	var ch chan int  // ch == nil

	select {
	case v, ok := <-ch:  // 读写都会阻塞。
		fmt.Println(v, ok)
	default:
		fmt.Println("default") // default
	}
}
