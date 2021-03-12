package main

import "fmt"

// 下面这段代码输出什么？
func main() {
	// a =  [1 12 13 4 5], r =  [1 12 13 4 5]
	// a := []int{1, 2, 3, 4, 5}

	// a =  [1 12 13 4 5], r =  [1 2 3 4 5]
	a := [...]int{1, 2, 3, 4, 5}
	var r [5]int

	// 切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，
	// 所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素。
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("a = ", a)
	fmt.Println("r = ", r)
}
