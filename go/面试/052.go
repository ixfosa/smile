package main

import "fmt"

const (
	a = iota
	b = iota
)

const (
	c = "ixfosa"
	d = iota
	e = iota
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
// 0
// 1
// ixfosa
// 1
// 2

// iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。
// iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。