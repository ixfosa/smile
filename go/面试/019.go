package main

import "fmt"

// 下面这段代码输出什么？
	// A.13.1 //
	// B.13
	// C.compilation error

// 参考答案及解析：C。a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错。
func main() {
	a := 5
	b := 8.1
	// fmt.Println(a + b) //  invalid operation: a + b (mismatched types int and float64)
	fmt.Println(a + int(b)) // 13
}
