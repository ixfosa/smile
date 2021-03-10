package main

import "fmt"

// 下面这段代码输出什么？
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2

	defer calc("1", a, calc("10", a, b))

	a = 0

	defer calc("2", a, calc("20", a, b))

	b = 1
}

// 会先执行 calc() 函数的 b 参数，即：calc(“10”,a,b)，输出：10 1 2 3

// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4

