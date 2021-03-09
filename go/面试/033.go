package main

import "fmt"

// 下面这段代码输出什么？
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	w := Work{3}

	var a A = w
	var b B = w

	fmt.Println(a.ShowA()) // 13
	fmt.Println(b.ShowB()) // 23

	// t := Work{3}
	// var a1 A = t
	// var b1 B = t
	// a1.ShowB()
	// b1.ShowA()
	// 接口的静态类型。
	// a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；
	// a 的静态类型是 A，b 的静态类型是 B，
	// 接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。
	// a.ShowB undefined (type A has no field or method ShowB)
	// b.ShowA undefined (type B has no field or method ShowA)

	// 类型断言
	var t A = Work{3}
	s := t.(Work)
	fmt.Println(s.ShowA())  // 13
	fmt.Println(s.ShowB())  // 23


}
// 知识点：接口。一种类型实现多个接口，
// 结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。