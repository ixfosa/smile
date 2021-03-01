package main

import "fmt"

// 下面这段代码输出什么，说明原因。
func main() {
	slice := []int{1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice{
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k,"->",*v)
	}

	// 正确的写法:
	for key, val := range slice{
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k,"->",*v)
	}
}

// res:
	// 0 -> 3
	// 1 -> 3
	// 2 -> 3

// for range 循环的时候会创建每个元素的副本，而不是元素的引用，
// 所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，
// 因为最后 val 被赋值为3，所有输出都是3.

