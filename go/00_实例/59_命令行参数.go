package main

import (
	"fmt"
	"os"
)

// 命令行参数是指定程序运行参数的一个常见方式。
// 例如，go run hello.go，程序 go 使用了 run 和 hello.go 两个参数。

// os.Args 提供原始命令行参数访问功能。
// 注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
/*
	D:\code\GoProjects\src\smile\00_实例>go run 59_命令行参数.go hello go !
	[C:\Users\86138\AppData\Local\Temp\go-build406843174\command-line-arguments\_obj\exe\59_命令行参数.exe hello go !]
	[hello go !]
	!
 */