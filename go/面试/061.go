package main

import "fmt"

// 下面这段代码输出什么？
func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b [5]int

	// range 表达式是副本参与循环，就是说参与循环的是 a 的副本，而不是真正的 a。
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		b[i] = v
	}
	fmt.Println("a = ", a) // [1, 12, 13, 4, 5]
	fmt.Println("b = ", b) // [1, 2, 3, 4, 5]


	// 如果想要 b 和 a 一样输出，修复办法：
	// 使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，
	// 因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。
	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		b[i] = v
	}
	fmt.Println("a = ", a) // [1, 12, 13, 4, 5]
	fmt.Println("b = ", b) // [1, 12, 13, 4, 5]
}
