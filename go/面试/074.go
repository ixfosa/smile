package main

import "fmt"

/*
关于无缓冲和有冲突的channel，下面说法正确的是？
	A. 无缓冲的channel是默认的缓冲为1的channel；
	B. 无缓冲的channel和有缓冲的channel都是同步的；
	C. 无缓冲的channel和有缓冲的channel都是非同步的；
	D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；
// 参考答案及解析：D。

 */


// 下面代码是否能编译通过？如果通过，输出什么？

// interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
// 这里的 x 的动态类型是 *int，所以 x 不为 nil。
func Foo74(x interface{})  {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo74(x)  // non-empty interface

	var z *interface{} = nil
	Foo74(z) // non-empty interface

	var y *interface{}
	Foo74(y) // non-empty interface

	var p interface{}
	Foo74(p) // empty interface
}
