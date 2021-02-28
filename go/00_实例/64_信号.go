package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go 通过向一个通道发送 os.Signal 值来进行信号通知。
	// 我们将创建一个通道来接收这些通知（同时还创建一个用于在程序可以结束时进行通知的通道）。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify 注册这个给定的通道用于接收特定信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
