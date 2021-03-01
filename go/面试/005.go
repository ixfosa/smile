package main

import "fmt"

// new() 与 make() 的区别
//参考答案：
	// new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

	// new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。
	// 换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

	// make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，
	// 是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
func main() {
	mapPtr := new(map[int]int)
	fmt.Printf("mapPtr: %p \n", mapPtr) // mapPtr: 0xc042004028
	// (*(mapPtr))[1] = 1   // panic: assignment to entry in nil map
	fmt.Printf("val: %v \n", *mapPtr)

	m := make(map[int]int)
	fmt.Printf("mPtr: %p \n", &m) // mPtr: 0xc042004038

	m[2] = 2
	fmt.Println("m: ", m) // m:  map[2:2]

	mapPtr = &m
	fmt.Printf("mapPtr: %p \n", mapPtr) // mPtr: 0xc042004038
	fmt.Printf("val: %v \n", *mapPtr) // val: map[2:2]
}
