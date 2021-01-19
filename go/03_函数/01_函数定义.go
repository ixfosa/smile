package main

import (
	"fmt"
)

/******************golang函数特点：*********************/
/*
	• 无需声明原型。
    • 支持不定 变参。
    • 支持多返回值。
    • 支持命名返回参数。
    • 支持匿名函数和闭包。
    • 函数也是一种类型，一个函数可以赋值给变量。

    • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
    • 不支持 重载 (overload)
    • 不支持 默认参数 (default parameter)。
 */


/*******************函数声明：********************//*
函数声明包含一个 函数名，参数列表， 返回值列表 和 函数体。
如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。

函数可以没有参数或接受多个参数。
	*注意类型在变量名之后 。

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

函数可以返回任意数量的返回值。

使用关键字 func 定义函数，左大括号依旧不能另起一行。
 */

//类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
func test1(x, y int, s string) (int, string)  {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func test2(fn func() int) int {
	return fn()
}

type FormatFunc func (s string, x, y int) string

func format (fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main()  {
	s1 := test2(func() int {
		return 99
	})

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println("s1 : ", s1) //s1 :  99
	fmt.Println("s2 : ", s2) //s2 :  10, 20
}

/***************************************//*

有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符。

    package math
    func Sin(x float64) float //implemented in assembly language

 */


