package main

import "fmt"

// 下面代码输出什么？
	//- A. 4
	//- B. compilation error    T
type Math struct {
	x, y int
}

var m055 = map[string]Math {
	"foo": Math{1, 2},
}

var m0551 = map[string]*Math {
	"foo": &Math{1, 2},
}

func main() {
	// m055["foo"].x = 4   // cannot assign to struct field m["foo"].x in map

	// 有两个解决办法：
	// 1.使用临时变量
	tmp := m055["foo"]
	tmp.x = 4
	fmt.Println(m055["foo"]) // {1 2}

	m055["foo"] = tmp
	fmt.Println(m055["foo"].x) // 4

	// 2.修改数据结构
	/*
	var m0551 = map[string]*Math {
		"foo": &Math{1, 2},
	}
	 */
	m0551["foo"].x = 4
	fmt.Println(m0551["foo"]) // &{4 2}
}
// 错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，
// 但 go 中的 map 的 value 本身是不可寻址的。