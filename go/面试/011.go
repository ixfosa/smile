package main

import "fmt"

// 2.下面这段代码能否编译通过？如果可以，输出什么？


const (
	x = iota  // 0
	_         // 可跳过的值
	y         // 2
	z = "zz"  // "zz"   中间插队
	k		  // "zz"
	p = iota  // 5
)

const (
	n = iota  // 0
	m		  // 1
	i = iota  // 2
	j         // 3
)

// iota只能在常量的表达式中使用。
// 每次 const 出现时，都会让 iota 初始化为0.【自增长】

func main()  {
	fmt.Println(x,y,z,k,p)
	fmt.Println(n, m, i, j)
}

// 参考答案及解析：编译通过，输出：0 2 zz zz 5。知识点：iota 的使用。