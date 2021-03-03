package main

import (
	"fmt"
)

// 1.定义一个包内全局字符串变量，下面语法正确的是（）
	// A. var str string
	// B. str := “”
	// C. str = “”
	// D. var str = “”
// 参考答案及解析：AD。B 只支持局部变量声明；C 是赋值，str 必须在这之前已经声明；



// 下面这段代码输出什么?
func hello271(i int)  {
	fmt.Println(i)

}
func main() {
	i := 1
	defer hello271(i) // 1

	defer func() {
		fmt.Println(i) // 10
	}()

	defer func(i int) {
		fmt.Println(i) // 1
	}(i)

	i += 9
}
// 参考答案及解析：1。
	// 这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 1.