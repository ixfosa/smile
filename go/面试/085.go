package main

import "fmt"

// 下面代码有几处错误的地方？请说明原因。
	// 有 1 出错误，不能对 nil 的 map 直接赋值，需要使用 make() 初始化。但可以使用 append() 函数对为 nil 的 slice 增加元素。
func main() {
	var s []int
	s = append(s, 1)

	var m map[string]int
	// 修复
	m = make(map[string]int)
	m["A"] = 1      // panic: assignment to entry in nil map

	fmt.Println(s, m)
}
