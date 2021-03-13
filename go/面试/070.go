package main

import "fmt"

/*
关于channel的特性，下面说法正确的是？
	A. 给一个 nil channel 发送数据，造成永远阻塞
	B. 从一个 nil channel 接收数据，造成永远阻塞
	C. 给一个已经关闭的 channel 发送数据，引起 panic
	D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
// 参考答案及解析：ABCD。



 */

// 下面代码有什么问题？
const i070 = 100
var j070 = 123

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("ch is nil") // ch is nil
	}

	// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
	// fmt.Println(&i070, i070) // cannot take the address of i070

	fmt.Println(&j070, j070) // 0x52a1e8 123

}
