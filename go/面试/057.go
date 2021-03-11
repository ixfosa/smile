package main

import (
	"fmt"
	"log"
)

// 下面这段代码输出什么？
	// A. 5 5
	// B. runtime error

var p057 *int
func foo() (*int, error)  {
	var i int = 5
	return &i, nil
}

func bar()  {
	fmt.Println(*p057)
}
func main() {
	// p 是新定义的变量，会遮住全局变量 p
	p057, err := foo()

	if err != nil {
		log.Panic(err)
		return
	}

	// 执行到bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash。
	bar() // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(*p057)
}

// 变量作用域。
// 问题出在操作符:=，对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。



