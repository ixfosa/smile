package main

import "fmt"

// 下面代码输出什么？
	// A. hello
	// B. xello
	// C. compilation error

// 参考代码及解析：C。知识点：常量，Go 语言中的字符串是只读的。

func main() {
	str := "hello"
	// str[0] = "h" 	// cannot assign to str[0]
	fmt.Println(str)

	strs := []rune(str)  // []byte(str)
	fmt.Printf("%T\n", strs) // []int32
	strs[0] = 'h'
	str = string(strs)
	fmt.Println( str) // hello
}
