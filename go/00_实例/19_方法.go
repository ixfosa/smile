package main

import "fmt"

type rect struct {
	h, w int
}

func (r *rect) area() int {
	return r.h * r.w
}

func (r rect) perim() int {
	return 2 * (r.h + r.w)
}

func main() {
	r1 := rect{h: 3, w: 9}

	fmt.Println(r1.area())  // 27
	fmt.Println(r1.perim()) // 24

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据
	r2 := &r1
	fmt.Println(r2.area())  // 27
	fmt.Println(r2.perim()) // 24
}
