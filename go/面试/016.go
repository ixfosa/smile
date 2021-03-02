package main

import "fmt"

// 关于channel，下面语法正确的是()
	// A. var ch chan int
	// B. ch := make(chan int)
	// C. <- ch
	// D. ch <-

// 参考答案及解析：ABC。A、B都是声明 channel；C 读取 channel；写 channel 是必须带上值，所以 D 错误。
func main() {
	var ch chan int

	fmt.Println(ch) // <nil>

	ch = make(chan int)

	fmt.Println(ch) // 0xc04200c0c0
}
