package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 从一个简单的命令开始，打印一些信息到标准输出流。
	// exec.Command 函数帮助我们创建一个表示这个外部进程的对象。
	// func Command(name string, arg ...string) *Cmd
	dataCmd := exec.Command("echo", "hello")
	dataOut, err := dataCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dataOut))
}
