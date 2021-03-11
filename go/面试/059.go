package main

import (
	"fmt"
	"time"
)

// 下面这段代码输出什么？为什么
func main() {
	var m = [...]int{1, 2, 3}

	// for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。
	// 各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，而不是各个goroutine启动时的i, v值。
	// 可以理解为闭包引用，使用的是上下文环境的值。
	/*
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}
	// 2 3
	// 2 3
	// 2 3
	 */
	for i, v := range m {
		go func(int, int) {   // 形参未命名， 乱序
			fmt.Println(i, v)
		}(i, v)
	}


	/*
	// 1.使用函数传递
	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	// 0 1
	// 1 2
	// 2 3
	 */

	/*
	// 2.使用临时变量保留当前值
	for i, v := range m {
		i := i     //  // 这里的 := 会重新声明变量，而不是重用
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
	// 0 1
	// 1 2
	// 2 3
	 */



	time.Sleep(time.Second * 3)
}
