package main

import "fmt"

func main() {
	// var 声明 1 个或者多个变量
	var a string = "fo"
	fmt.Println("a = ", a)

	// 一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println("a = ", b, "b = ", c)

	// 自动推断已经初始化的变量类型
	var d = true
	fmt.Println("d = ", d) //d =  true

	// 声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
	var e int
	fmt.Println("e = ", e) // e =  0

	// := 语句是申明并初始化变量的简写
	f := "fo"
	fmt.Println("f = ", f)
}
