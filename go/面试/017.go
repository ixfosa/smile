package main

import "fmt"

// 下面这段代码输出什么？
	// 打印一个 map 中不存在的值时，返回元素类型的零值。
	// 这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0。
type person struct {
	name string
}

func main() {
	// m := make(map[person]int)
	var m map[person]int

	fmt.Println(m) // map[]

	p := person{"Zi Qing" }
	// m[p] = 1   panic: assignment to entry in nil map

	fmt.Println(m[p]) // 0
}
