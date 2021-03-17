package main

import "fmt"

// 下面这段代码输出什么？
type T struct {
	ls []int
}

func change104(t T)  {
	t.ls[0] = 100
}

// 调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针。
func main() {
	t := T{
		ls: []int{1, 2, 3,},
	}
	change104(t)
	fmt.Println(t.ls[0]) // 100
}


