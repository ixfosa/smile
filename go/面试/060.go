package main

import "fmt"

// 下面这段代码输出什么？
func f60(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()

	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f60(3)) // 7
}

// 第一步执行 r = n +1 = 4，
// 接着执行第二个 defer，由于此时 f() 未定义，引发异常，
// 随即执行第一个 defer，异常被 recover()，r = r + n = 4 +3 = 7 程序正常执行，最后 return。

