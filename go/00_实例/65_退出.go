package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用 os.Exit 时 defer 将不会 执行，所以这里的 fmt.Println将永远不会被调用。
	defer fmt.Println("!")

	// 退出并且退出状态为 3。
	os.Exit(3)

	// 注意我们程序中的 ! 永远不会被打印出来
}
