package main


// 	面代码有什么问题？
func main() {
	x := int{
		1,
		// 2     // syntax error: unexpected newline, expecting comma or } 未加逗号
		// 用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。
	}
	x = x

	y := []int{3, 4}

	y = y
}
