package main

// 下面的代码有几处语法问题，各是什么？
/*
func main() {
	var s string = nil

	if s == nil {
		s = "fo"
	}
	fmt.Println(s)
}
 */
// 两个地方有语法问题。golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。