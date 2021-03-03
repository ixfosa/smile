package main

import "fmt"

// 关于 cap() 函数的适用类型，下面说法正确的是()
//	A. array
//	B. slice
//	C. map
//	D. channel

//参考答案及解析：ABD。知识点：cap()，cap() 函数不适用 map。
func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	// m := map[int]int{1: 1}
	ch := make(chan int, 3)

	fmt.Println(cap(arr)) // 3
	fmt.Println(cap(slice)) // 3
	// fmt.Println(cap(m)) // invalid argument m (type map[int]int) for cap
	fmt.Println(cap(ch)) // 3
}
