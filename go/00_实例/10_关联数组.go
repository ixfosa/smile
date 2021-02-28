package main

import (
	"fmt"
)

// map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）。
func main() {
	// 要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type).
	m := make(map[string]int)

	//  make[key] = val 语法来设置键值对
	m["a"] = 0
	m["b"] = 1
	m["c"] = 2

	// Println 来打印一个 map 将会输出所有的键值对。
	fmt.Println("m = ", m) // m =  map[a:0 b:1 c:2]

	// 使用 name[key] 来获取一个键的值
	fmt.Println("m[\"b\"] = ", m["b"]) // m["b"] =  1

	// map 调用内建的 len 时，返回的是键值对数目
	fmt.Println("len = ", len(m))

	// 内建的 delete 可以从一个 map 中移除键值对
	delete(m, "c")
	fmt.Println("m = ", m) // m =  map[a:0 b:1]
	fmt.Println("len = ", len(m)) // len =  2

	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。
	// 这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	v1, ok1 := m["a"]
	if ok1 {
		fmt.Println("v1 = ", v1) // v1 =  0
	}

	_, ok2 := m["c"]
	if !ok2 {
		fmt.Println("m[\"c\"] 不存在") // m["c"] 不存在
	}

	// 在同一行申明和初始化一个新的map。
	m2 := map[string]int{"d": 3, "e": 4}
	fmt.Println("m2 = ", m2) //m2 =  map[d:3 e:4]
}
