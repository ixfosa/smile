package main

import "fmt"

// 下面代码输出什么？

type T107 struct {
	n int
}

func main() {
	m := make(map[int]T107)

	// map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。
	// m[0].n = 1 // cannot assign to struct field m[0].n in map

	fmt.Println(m[0] , 1)
}
