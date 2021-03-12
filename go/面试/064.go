package main

import "fmt"

// 下面这段代码输出结果正确吗？
	// 输出：
	// {A} {B} {C}
	// &{A} &{B} &{C}
// s2 的输出结果错误。s2 的输出是 &{C} &{C} &{C}

// for range 使用短变量声明(:=)的形式迭代变量时，变量 i、value 在每次循环体中都会被重用，而不是重新声明。
// 所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。
// 因此，s2 输出的时候显示出了三个 &{c}。
type Foo struct {
	bar string
}
func main() {
	s1 := []Foo{
		{"A",},
		{"B",},
		{"C",},
	}

	s2 := make([]*Foo, len(s1))

	for i, v := range s1 {
		s2[i] = &v

		// 可行的解决办法如下：
		// s2[i] = &s1[i]
	}

	fmt.Println(s1[0], s1[1], s1[2]) // {A} {B} {C}
	fmt.Println(s2[0], s2[1], s2[2]) // &{C} &{C} &{C}
}
