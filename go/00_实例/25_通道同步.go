package main

import (
	"fmt"
	"time"
)

// 可以使用通道来同步 Go 协程间的执行状态。这里是一个使用阻塞的接受方式来等待一个 Go 协程的运行结束
func worker(done chan string)  {
	fmt.Println("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done...")
	done <- "ok"
}
func main() {
	done := make(chan string, 1)
	go worker(done)

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	// 如果把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
	<- done
}
