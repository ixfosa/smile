package main

import "fmt"

// 下面代码有什么问题？

func main() {
	m := make(map[string]int, 2)
	// cap(m)
	fmt.Println(len(m))
	// 1. 使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
	// 2. cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。
}
