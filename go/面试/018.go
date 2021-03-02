package main

import "fmt"

// 下面这段代码输出什么？
func hello181(num ...int)  {
	fmt.Printf("%T \n", num)  // []int
	num[0] = 99
}
func main() {
	ints := []int{1, 2, 3}
	hello181(ints...)
	fmt.Println(ints[0]) // 99
}
