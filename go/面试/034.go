package main

import "fmt"

// 切片 a、b、c 的长度和容量分别是多少？
func main() {
	s := [3]int{1, 2, 3}

	a := s[:0]	// 0, 3
	b := s[:2]  // 2, 3
	c := s[1:2:cap(s)] // 1, 2
	d := c[:1]  // 1, 2
	fmt.Println("a len: ", len(a), "-- cap: ", cap(a))
	fmt.Println("b len: ", len(b), "-- cap: ", cap(b))
	fmt.Println("c len: ", len(c), "-- cap: ", cap(c))
	fmt.Println("d len: ", len(d), "-- cap: ", cap(d))
}

// 截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。
// 在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
// 操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
