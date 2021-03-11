package main

import "fmt"

// 下面这段代码输出什么？为什么？
type People053 interface {
	Show()
}

type Student053 struct {
	
}

func (s *Student053) Show()  {
	
}
func main() {
	var s *Student053
	if s == nil {
		fmt.Println("s is nil") // s is nil
	} else {
		fmt.Println("s is not nil")
	}

	var p People053 = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil") // p is not nil
	}
}

// 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
