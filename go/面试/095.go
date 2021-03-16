package main

import "fmt"

// 下面代码最后一行输出什么？请说明原因。
func main() {
	x := 1
	fmt.Println(x) // 1

	{
		fmt.Println(x) // 1

		i, x := 6, 6
		fmt.Println(i, x) // 6, 6
	}
	// fmt.Println(i, x) // undefined: i

	fmt.Println(x) // 1
}
