package main

import (
	"fmt"
	"math"
)

// Go 支持字符、字符串、布尔和数值 常量 。
func main() {

	// const 语句可以出现在任何 var 语句可以出现的地方

	// const 用于声明一个常量
	const str string = "Zi Qing"

	// str = "fo" // cannot assign to str

	fmt.Println(str)


	// 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	// fmt.Println(int64(d))
	// 当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数。

	const n = 50000
	const d = 3e20 / n

	fmt.Printf("%T \n", n) // int
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
