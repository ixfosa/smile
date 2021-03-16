package main

import "fmt"

// 下面代码输出什么？

func text(x byte)  {
	fmt.Println(x)
}
func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	text(c) // 34
}
// 与 rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。