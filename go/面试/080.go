package main

// 请指出下面代码的错误？
	// 变量 one、two 和 three 声明未使用。
	// 知识点：未使用变量。如果有未使用的变量代码将编译失败。
	// 但也有例外，函数中声明的变量必须要使用，但可以有未使用的全局变量。函数的参数未使用也是可以的。

/*
var gvar int

func main() {
	var one int // one declared and not used
	two := 2 // two declared and not used
	var three int
	three = 3 // three declared and not used

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")

}
 */