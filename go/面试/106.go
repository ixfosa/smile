package main

// 下面的代码能否正确输出？
func main() {
	var f1 = func() {}
	var f2 = func() {}

	if f1 != f2 {}   // invalid operation: f1 != f2 (func can only be compared to nil) 函数只能与 nil 比较。
	println("fn1 not equal fn2")
}
