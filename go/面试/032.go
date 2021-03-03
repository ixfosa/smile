package main

import "fmt"


// 下面这段代码输出什么？
	// A. A
	// B. 65
	// C. compilation error

//参考答案及解析：A。UTF-8 编码中，十进制数字 65 对应的符号是 A。

func main() {
	i := 65
	fmt.Println(string(i)) // A
}
