package main

import (
	"fmt"
)

// 下面这段代码输出什么？

// Go 提供的语法糖…，可以将 slice 传进可变函数,不会创建新的切片
func change(s ...int) {
	// s ...int ptr: 0xc0420563e0 --- len: 5 --- cap: 5
	fmt.Printf("s ...int ptr: %p --- len: %d --- cap: %d \n", &s, len(s), cap(s))
	// append()操作。Go 提供的语法糖…，
	// append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；

	// s ptr: 0xc0420563e0 --- len: 6 --- cap: 10
	s = append(s, 3)

	// slice ptr: 0xc0420563a0 --- len: 5 --- cap: 5
	fmt.Printf("s ptr: %p --- len: %d --- cap: %d \n", &s, len(s), cap(s))
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	// fmt.Printf("slice ptr: %p --- len: %d --- cap: %d \n", &slice, len(slice), cap(slice))

	// change(slice...)
	// fmt.Println(slice) // [1, 2, 0, 0, 0]


	// [1, 2]
	s := slice[:2]
	fmt.Printf("slice[:2]... ptr: %p --- len: %d --- cap: %d \n", &s, len(slice[:2]), cap(slice[:2]))
	// 第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，
	// 它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，
	// 所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。
	change(s...)
	// slice[:2]... ptr: 0xc0420023e0 --- len: 2 --- cap: 5
	// s ...int ptr: 0xc042002420 --- len: 2 --- cap: 5
	// s ptr: 0xc042002420 --- len: 3 --- cap: 5
	// [1 2 3 0 0]

	fmt.Println(slice) // [1, 2, 3, 0, 0]


	c := make([]int, 3, 6)
	c[0] = 1
	c[1] = 2
	// c[:2] // len = 2, cap = 6
	cc := c[:2]
	// cc = append(cc, 3) // [1 2 3] [1 2 3]
	a := append(cc, 3) // [1 2 3] [1 2 3]

	fmt.Printf("%p \n", &cc)  // 0xc042056420

	fmt.Println(c, a)
	fmt.Printf("%p \n", &cc)

	a = append(c[1:], 3) // [1 2 0] [2 0
	fmt.Println(c, a)
}

/*
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
 */