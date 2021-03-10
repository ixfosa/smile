package main

import "fmt"

// 下面的两个切片声明中有什么区别？哪个更可取？
func main() {
	var a []int  // 声明的是 nil 切片；不会分配内存，优先选择。

	b := []int{}  // 声明的是长度和容量都为 0 的空切片。

	fmt.Println(a, b)
}
