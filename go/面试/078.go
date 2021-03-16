package main

import "fmt"

// 下面这段代码输出什么？
// 字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一
func main() {
	x := []int{2: 2, 3, 0: 1}
	fmt.Println(x) // [1, 0, 2, 3]
}
