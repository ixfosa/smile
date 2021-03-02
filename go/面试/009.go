package main

import "fmt"

// 下面这段代码能否通过编译？如果通过，输出什么？
	// 参考答案及解析：编译不通过，cannot use i (type int) as type MyInt1 in assignment。

type MyInt1 int  // 基于类型 int 创建了新类型 MyInt1
type  Myint2 = int // 创建了 int 的类型别名 MyInt2，注意类型别名的定义时 =

func main() {
	var i int = 1

	// 代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；
	// var i1 MyInt1 = i   // cannot use i (type int) as type MyInt1 in assignment

	//赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i).
	var i1 MyInt1 = MyInt1(i)

	// 而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	var i2 Myint2 = i

	fmt.Println(i1, i2) // 1 1
}
