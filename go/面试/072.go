package main

import "fmt"

/*
关于异常的触发，下面说法正确的是？
	A. 空指针解析；
	B. 下标越界；
	C. 除数为0；
	D. 调用panic函数；
// 参考答案及解析：ABCD。


 */

// 下面代码输出什么？
func main() {
	x := []string{"a", "b", "c"}

	for v := range x {
		fmt.Println(v) // 0, 1, 2
	}

	for _, v := range x {
		fmt.Println(v) // "a", "b", "c"
	}

}
