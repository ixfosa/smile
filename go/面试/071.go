package main

import "fmt"

// 下面代码能否编译通过？如果通过，输出什么？

func GetValue71(m map[int]string, id int) (string, bool) {
	if _, exist :=  m[id]; exist {
		return "exist", true
	}
	return nil, false //  cannot use nil as type string in return argument
}

func main() {
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	val, ok := GetValue71(m, 1)
	fmt.Printf(val, ok)
}

// nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
// 但是如果不特别指定的话，Go 语言不能识别类型，所以会报错: cannot use nil as type string in return argument.