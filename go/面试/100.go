package main

import "fmt"

// 下面代码有什么问题？
	// 编译错误：
type foo100 struct {
	bar100 int
}

func main() {
	var foo foo100
	// foo.bar100, f := 1, 2
	fmt.Println(foo.bar100, f) // non-name foo.bar100 on left side of :=
}
