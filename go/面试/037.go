package main

import "fmt"

func increaseA() int {
	var i int

	defer func() {
		i++
	}()

	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()

	return r
}

func main() {
	// increaseA() 的返回参数是匿名，increaseB() 是具名。
	fmt.Println(increaseA())  // 0
	fmt.Println(increaseB())  // 1
}
