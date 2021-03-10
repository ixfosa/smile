package main

import (
	"fmt"
	"time"
)

// 下面这段代码输出什么？
func main() {
	m := map[int]string{0: "zero", 1: "one"}

	for {
		for k, v := range m{
			fmt.Println(k, v)   // map 的输出是无序的。
		}
		time.Sleep(time.Second)
	}
}
	// 0 zero
	// 1 one
// 或
	// 1 one
	// 0 zero