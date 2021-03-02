package main

import "fmt"

// 通过指针变量 p 访问其成员变量 name，有哪几种方式？
	// A.p.name
	// B.(&p).name
	// C.(*p).name
	// D.p->name
// 参考答案及解析：AC ,   & 取址运算符，* 指针解引用。

func main() {
	sn1 := struct {
		name string
		age  int
	}{"Zi Qing", 18}

	p := &sn1

	fmt.Printf("%v \n", p)

	fmt.Printf("%v \n", p.name) // Zi Qing

	fmt.Printf("%v \n", (*p).name) // Zi Qing
}
