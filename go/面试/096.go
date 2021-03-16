package main

import "fmt"

// 下面代码有什么问题吗？
func main() {
	for i := 0; i < 10; i++ {
		// loop:
			fmt.Println(1)
	}
	// goto loop // goto loop jumps into block starting at .\096.go:7:26
}

// goto 不能跳转到其他函数或者内层代码。编译报错：