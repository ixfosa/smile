package main

import "fmt"

func main() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	panic("err")
	// defer func() { fmt.Println("4") }()
}



// 3
// 2
// 1
// panic: err

// 参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic