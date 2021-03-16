package main

import "fmt"

// 下面代码输出什么？

func main() {
	var a = []int{1, 2, 3, 4, 5}
	b := make([]int, 0)

	for i, v := range a {
		fmt.Printf("range a %p \n", &a)
		if i == 0 {
			a = append(a, 6, 7)
			fmt.Printf("if a %p \n", &a)
		}
		b = append(b, v)
	}
	fmt.Println(b) // [1 2 3 4 5]
}

// a 在 for range 过程中增加了两个元素

// len 由 5 增加到 7，但 for range 时会使用 a 的副本 a’ 参与循环，副本的 len 依旧是 5，
// 因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素。