package main

import "fmt"

// 常规的通过通道发送和接收数据是阻塞的。
// 然而，可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
func main() {
	// 这里是一个非阻塞接收的例子。
	// 如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。
	// 如果不是，就直接到 default 分支中。
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received") // no message received
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent") // no messag
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity") // no activity
	}
}
