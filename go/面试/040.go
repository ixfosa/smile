package main

import "fmt"

// 下面这段代码正确的输出是什么？
func f40()  {
	defer fmt.Println("D")
	fmt.Println("F")
}
func main() {
	f40()
	fmt.Println("H")
}
// F
// D
// H

