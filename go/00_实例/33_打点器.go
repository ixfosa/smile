package main

import (
	"fmt"
	"time"
)

// 定时器 是当你想要在未来某一刻执行一次时使用的 - 打点器 则是当你想要在固定的时间间隔重复执行准备的。

// 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
func main() {
	// 打点器和定时器的机制有点相似：一个通道用来发送数据。这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
	ticker := time.NewTicker(time.Second * 1)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// Tick at 2021-02-20 16:38:46.8183082 +0800 CST m=+1.004982601
// Tick at 2021-02-20 16:38:47.8173336 +0800 CST m=+2.004008001
// Tick at 2021-02-20 16:38:48.8179761 +0800 CST m=+3.004650501
// Tick at 2021-02-20 16:38:49.8179247 +0800 CST m=+4.004599101
// Ticker stopped