package main

import "fmt"

// 下面这段代码输出什么？

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	incr(&v)
	fmt.Println(v) // 2
}
