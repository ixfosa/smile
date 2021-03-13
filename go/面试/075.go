package main

import (
	"fmt"
	"time"
)

// 下面代码输出什么？
func main() {
	ch := make(chan int, 100)

	// A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// B
	go func() {
		for {
			a, ok := <- ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a = ", a)
		}
	}()
	// 当 A 协程还没起时，主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic
	close(ch) // panic: send on closed channel
	fmt.Println("OK!!!")
	time.Sleep(time.Second * 10)
}
