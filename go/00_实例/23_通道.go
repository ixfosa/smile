package main

import "fmt"

// 通道 是连接多个 Go 协程的管道。可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收。
func main() {
	// 使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
	msg := make(chan string)

	go func() {
		// 使用 channel <- 语法 发送 一个新的值到通道中。
		msg <- "Hello"
	}()

	// 使用 <-channel 语法从通道中 接收 一个值。
	res := <- msg
	fmt.Println(res) // Hello

	// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
}
