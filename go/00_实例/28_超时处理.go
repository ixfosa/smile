package main

import (
	"fmt"
	"time"
)

// 超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "ok1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("msg1 = ", msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1") // timeout 1
	}


	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "ok2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println("msg2 = ", msg2) // msg2 =  ok2
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3")
	}
}
