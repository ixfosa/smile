package main

import "fmt"

// 下面的代码有什么问题？
	// 编译可以通过。知识点：常量。常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。
func main() {
	const x = 19
	const y = 1.99
	fmt.Println(x)
}
