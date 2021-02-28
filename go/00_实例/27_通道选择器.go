package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "two"
	}()

	// 首先接收到值 "one"，然后就是预料中的 "two"
	for i := 0; i < 2; i++ {
		//使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
