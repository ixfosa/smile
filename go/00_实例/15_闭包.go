package main

import "fmt"

// Go 支持通过 闭包来使用 匿名函数。
func main() {
	nextInt := intSeq()
	// 调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。
	// 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
}

// intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏 变量 i
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

