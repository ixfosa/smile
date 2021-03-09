package main

import (
	"fmt"
)

// 下面代码中 A B 两处应该怎么修改才能顺利编译？
/*
func main() {
	var m map[string]int
	m["a"] = 1
	if v := m["b"]; v != nil {  // cannot convert nil to type int
		fmt.Println(v)
	}
}



func main() {
	var m map[string]int  // panic: assignment to entry in nil map, 未初始化
	m["a"] = 1
	if v := m["b"]; v != 0 {  // cannot convert nil to type int
		fmt.Println(v)
	}
}
*/


func main() {
	var m map[string]int  // panic: assignment to entry in nil map, 未初始化
	m = make(map[string]int)
	m["a"] = 1
	if v := m["b"]; v != 0 {  // cannot convert nil to type int
		fmt.Println(v)
	}
}

// 在 A 处只声明了map m ,并没有分配内存空间，不能直接赋值，需要使用 make()，都提倡使用 make() 或者字面量的方式直接初始化 map。
// B 处，v,k := m[“b”] 当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，k 返回 false。
