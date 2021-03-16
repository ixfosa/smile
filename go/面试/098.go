package main

import "fmt"

// 关于 slice 或 map 操作，下面正确的是？
/*
	A.
		var s []int
		s = append(s,1)

	B.
		var m map[string]int
		m["one"] = 1

	C.
		var s []int
		s = make([]int, 0)
		s = append(s,1)
	D.
		var m map[string]int
		m = make(map[string]int)
		m["one"] = 1

// ACD
*/

// 下面代码输出什么？

func test98(x int) (func(), func()) {
	return func() {
		fmt.Println(x)
		x += 10
	}, func() {
			fmt.Println(x)
		}
}

func main() {
	a, b := test98(10)  // 闭包引用相同变量。
	a() // 10
	b() // 20
}
