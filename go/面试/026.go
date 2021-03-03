package main

import "fmt"

// 下面这段代码输出什么？

type People struct {}

func (p *People) ShowA()  {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB()  {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
func main() {
	t := Teacher{}
	t.ShowB() // teacher showB

	t.ShowA() // showA, showB
	// Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法，
	// 需要注意的是代码调用的是 People 自己的 ShowB 方法。
}

// 参考答案及解析：teacher showB。
	// 知识点：结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；
	// 通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。
	// 此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。
	// 这个例子中的 ShowB() 就是同名方法。