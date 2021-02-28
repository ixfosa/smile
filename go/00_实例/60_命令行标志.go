package main

import (
	"flag"
	"fmt"
)

// Go 提供了一个 flag 包，支持基本的命令行标志解析。
func main() {
	// 基本的标记声明仅支持字符串、整数和布尔值选项。

	// 这里我们声明一个默认值为 "foo" 的字符串标志 word并带有一个简短的描述。
	// 这里的 flag.String 函数返回一个字符串指针（不是一个字符串值），在下面我们会看到是如何使用这个指针的。
	wordPtr := flag.String("word", "hello", "a string")

	// 使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	numPtr := flag.Int("num", 66, "an int")
	bolPtr := flag.Bool("fork", true, "a bool")

	// 使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	var svar string
	flag.StringVar(&svar, "svar", "world", "a string")

	// 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
	flag.Parse()

	// 将仅输出解析的选项以及后面的位置参数。注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用，从而得到选项的实际值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *bolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
/*
D:\code\GoProjects\src\smile\00_实例>go run 60_命令行标志.go -word = ixfosa
word: =
numb: 66
fork: true
svar: world
tail: [ixfosa]

D:\code\GoProjects\src\smile\00_实例>go run 60_命令行标志.go -word ixfosa
word: ixfosa
numb: 66
fork: true
svar: world
tail: []


D:\code\GoProjects\src\smile\00_实例>go run 60_命令行标志.go -word ixfosa -num 99
word: ixfosa
numb: 99
fork: true
svar: world
tail: []


D:\code\GoProjects\src\smile\00_实例>go run 60_命令行标志.go -word "Zi Qing" hahah hahah
word: Zi Qing
numb: 66
fork: true
svar: world
tail: [hahah hahah]


D:\code\GoProjects\src\smile\00_实例>go run 60_命令行标志.go -h
Usage of C:\Users\86138\AppData\Local\Temp\go-build938573566\command-line-arguments\_obj\exe\60_命令行标志.exe:
  -fork
        a bool (default true)
  -num int
        an int (default 66)
  -svar string
        a string (default "world")
  -word string
        a string (default "hello")





 */