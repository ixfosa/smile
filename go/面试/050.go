package main

import "fmt"

// 下面这段代码输出什么？为什么？
	// A. 1
	// B. compilation error

/*
func (i int) PrintInt() {
	fmt.Println(i)
}
func main() {
	var i int = 1
	i.PrintInt()
}
 */

// B。基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，
// 由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。

type myInt int
func (i myInt) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i myInt = 1
	i.PrintInt()
}