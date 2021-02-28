package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i) // initial:  1

	zeroval(1)
	fmt.Println("zeroval: ", i) // zeroval:  1

	//  &i 语法来取得 i 的内存地址，例如一个变量i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr: ", i) // zeroptr:  0
}
